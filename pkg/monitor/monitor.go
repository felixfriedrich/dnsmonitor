package monitor

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/dns"
	"dnsmonitor/pkg/model"
	"dnsmonitor/pkg/store"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// Monitor is an interface, which is used to enforce the use of CreateMonitor as the struct monitor can not be created
// differently, because it's private. It's also used to generate mocks.
//counterfeiter:generate . Monitor
type Monitor interface {
	Domains() []*model.Domain
	Config() configuration.Check
	Observe()
	Check()
	Run(interval int, silent bool)
}

// Monitor holds a Check and a Domain object from the store
type monitor struct {
	domains  []*model.Domain
	config   configuration.Check
	alerting alerting.API
	dns      dns.Interface
	mail     alerting.Mail
	ticker   *time.Ticker
}

// Domains returns a list of pointers to the Domains
func (m monitor) Domains() []*model.Domain {
	return m.domains
}

func (m monitor) Config() configuration.Check {
	return m.config
}

// CreateMonitor creates a Monitor fetching a domain from the store
func CreateMonitor(config configuration.Check, mail alerting.Mail, alerting alerting.API, dns dns.Interface) (Monitor, error) {
	domains := []*model.Domain{}
	for _, d := range config.Domains {
		d, err := store.Get(d)
		if err != nil {
			return nil, err
		}
		domains = append(domains, d)
	}
	m := monitor{domains: domains, config: config, alerting: alerting, dns: dns, mail: mail}
	return m, nil
}

// Check does a DNS query (for each domain) to find all answers until hitting A records and saves them in the store
func (m monitor) Check() {
	m.Observe()

	for _, d := range m.Domains() {
		diff, _ := d.GetDiff()

		if m.config.Mail && diff != "" {
			err := m.mail.Send(diff)
			if err != nil {
				log.Error(err)
			}
		}

		if m.config.SMS && diff != "" {
			m.alerting.SendSMS(diff)
		}

		err := store.Save(d)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Observe queries DNS and creates a Record of observed answers
func (m monitor) Observe() {

	for _, d := range m.domains {
		answers, err := m.dns.Query(d.Name, m.config.DNS)
		if err != nil {
			log.Error(err)
		}
		record := model.CreateRecord(answers)
		d.Observations = append(d.Observations, *record)
	}
}

func (m monitor) Run(interval int, silent bool) {

	// Tickers only execute their job after the interval time has passed once.
	// As the user might expect otherwise, the initial check is performed once before the Ticker is started.
	m.Check()
	m.printOutput(silent)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			m.Check()
			m.printOutput(silent)
		}
	}
}

func (m monitor) printOutput(silent bool) {
	for _, d := range m.Domains() {
		if !silent {
			fmt.Println("Checking domain", d.Name)
		}

		fmt.Println("Found", len(d.LastObservation().GetAnswers()), "answer(s).")
		for _, answers := range d.LastObservation().GetAnswers() {
			fmt.Println(answers)
		}

		diff, _ := d.GetDiff()
		fmt.Println(diff)
	}
}
