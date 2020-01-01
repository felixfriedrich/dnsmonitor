package monitor

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/dns"
	"dnsmonitor/pkg/model"
	"dnsmonitor/pkg/store"

	log "github.com/sirupsen/logrus"
)

// Monitor is an interface, which is used to enforce the use of CreateMonitor as the struct monitor can not be created
// differently, because it's private. It's also used to generate mocks.
//counterfeiter:generate . Monitor
type Monitor interface {
	Domain() *model.Domain
	Config() configuration.Config
	Observe() model.Record
	Check() model.Record
}

// Monitor holds a Config and a Domain object from the store
type monitor struct {
	domain   *model.Domain
	config   configuration.Config
	alerting alerting.API
	dns      dns.Interface
	mail     alerting.Mail
}

// Domain returns a pointer to the Domain
func (m monitor) Domain() *model.Domain {
	return m.domain
}

func (m monitor) Config() configuration.Config {
	return m.config
}

// CreateMonitor creates a Monitor fetching a domain from the store
func CreateMonitor(domain string, config configuration.Config, mail alerting.Mail, alerting alerting.API, dns dns.Interface) (Monitor, error) {
	d, err := store.Get(domain)
	if err != nil {
		return monitor{}, err
	}
	m := monitor{domain: d, config: config, alerting: alerting, dns: dns, mail: mail}
	return m, nil
}

// Check does a DNS query to find all answers until hitting A records and saves them in the store
func (m monitor) Check() model.Record {
	record := m.Observe()

	diff, _ := m.Domain().GetDiff()

	if m.config.Mail && diff != "" {
		err := m.mail.Send(diff)
		if err != nil {
			log.Error(err)
		}
	}

	if m.config.SMS && diff != "" {
		m.alerting.SendSMS(diff)
	}

	err := store.Save(m.domain)
	if err != nil {
		log.Fatal(err)
	}
	return record
}

// Observe queries DNS and creates a Record of observed answers
func (m monitor) Observe() model.Record {

	answers, err := m.dns.Query(m.domain.Name, m.config.DNS)
	if err != nil {
		log.Error(err)
	}

	record := model.CreateRecord(answers)
	m.domain.Observations = append(m.domain.Observations, *record)
	return *record
}
