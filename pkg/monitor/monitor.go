package monitor

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/dns"
	"dnsmonitor/pkg/model"
	"dnsmonitor/pkg/store"
	"strings"

	dnslib "github.com/miekg/dns"
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
	domain   *model.Domain
	config   config.Config
	alerting alerting.API
	dns      dns.Interface
}

// Domain returns a pointer to the Domain
func (m monitor) Domain() *model.Domain {
	return m.domain
}

func (m monitor) Config() config.Config {
	return m.config
}

// CreateMonitor creates a Monitor fetching a domain from the store
func CreateMonitor(domain string, config config.Config, alerting alerting.API, dns dns.Interface) (Monitor, error) {
	d, err := store.Get(domain)
	if err != nil {
		return monitor{}, err
	}
	m := monitor{domain: d, config: config, alerting: alerting, dns: dns}
	return m, nil
}

// Check does a DNS query to find all answers until hitting A records and saves them in the store
func (m monitor) Check() model.Record {
	record := m.Observe()

	diff, _ := m.Domain().GetDiff()

	if m.config.Mail {
		err := alerting.SendMail(diff)
		if err != nil {
			log.Error(err)
		}
	}

	if m.config.SMS {
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
	msg := dnslib.Msg{}
	msg.SetQuestion(m.domain.Name+".", dnslib.TypeA)

	r, _, err := m.dns.Exchange(&msg, m.config.DNS+":53")
	if err != nil {
		log.Fatal(err)
	}
	if r == nil {
		log.Fatal("dns Exchange return nil value")
	}

	answers := []string{}
	for _, a := range r.Answer {
		answers = append(answers, strings.Fields(a.String())[4])
	}

	record := model.CreateRecord(answers)
	m.domain.Observations = append(m.domain.Observations, *record)
	return *record
}
