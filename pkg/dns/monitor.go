package dns

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/model"
	"dnsmonitor/pkg/store"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
)

// Monitor is an interface, which is used to enforce the use of CreateMonitor as the struct monitor can not be created
// differently, because it's private
type Monitor interface {
	Domain() *model.Domain
	Config() config.Config
	Observe() model.Record
	Check() model.Record
}

// Monitor holds a Config and a Domain object from the store
type monitor struct {
	domain *model.Domain
	config config.Config
}

// Domain returns a pointer to the Domain
func (m monitor) Domain() *model.Domain {
	return m.domain
}

func (m monitor) Config() config.Config {
	return m.config
}

// CreateMonitor creates a Monitor fetching a domain from the store
func CreateMonitor(domain string, config config.Config) (Monitor, error) {
	d, err := store.Get(domain)
	if err != nil {
		return monitor{}, err
	}
	m := monitor{domain: d, config: config}
	return m, nil
}

// Check does a DNS query to find all answers until hitting A records and saves them in the store
func (m monitor) Check() model.Record {
	record := m.Observe()

	if m.config.Mail {
		diff, _ := m.Domain().GetDiff(record)
		err := m.sendMail(diff)
		if err != nil {
			log.Error(err)
		}
	}

	err := store.Save(m.domain)
	if err != nil {
		log.Fatal(err)
	}
	return record
}

func (m monitor) sendMail(diff string) error {
	c := config.CreateMailConfigFromEnvOrDie()
	if diff != "" {
		// Set up authentication information.
		auth := smtp.PlainAuth(
			"",
			c.Username,
			c.Password,
			c.Host,
		)
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		err := smtp.SendMail(
			c.Host+":"+strconv.Itoa(c.Port),
			auth,
			c.From,
			[]string{c.To},
			[]byte(diff),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// Observe queries DNS and creates a Record of observed answers
func (m monitor) Observe() model.Record {
	msg := dns.Msg{}
	msg.SetQuestion(m.domain.Name+".", dns.TypeA)
	dnsClient := dns.Client{}
	r, _, err := dnsClient.Exchange(&msg, m.config.DNS+":53")
	if err != nil {
		log.Fatal(err)
	}

	answers := []string{}
	for _, a := range r.Answer {
		answers = append(answers, strings.Fields(a.String())[4])
	}

	record := model.CreateRecord(answers)
	m.domain.Observations = append(m.domain.Observations, *record)
	return *record
}
