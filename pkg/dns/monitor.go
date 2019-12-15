package dns

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/store"
	"fmt"
	"github.com/miekg/dns"
	"github.com/pmezard/go-difflib/difflib"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Monitor holds all flags and a domain object from the store
type Monitor struct {
	flags  config.Flags
	domain store.Domain
}

func (m Monitor) Domain() store.Domain {
	return m.domain
}

// CreateMonitor creates a Monitor fetching a domain from the store
func CreateMonitor(domain string, flags config.Flags) (Monitor, error) {
	d, err := store.Get(domain)
	if err != nil {
		return Monitor{}, err
	}
	m := Monitor{flags: flags, domain: d}
	return m, nil
}

// Check does a DNS query to find all answers until hitting A records and saves them in the store
func (m Monitor) Check() {
	r := m.Observe()
	if !m.flags.Silent {
		fmt.Println("Found", len(r.GetAnswers()), "answer(s).")
		for _, aa := range r.GetAnswers() {
			fmt.Println(aa)
		}
	}

	diff, err := m.getDiff(r.GetAnswers())
	if err != nil {
		log.Error(err)
	}
	if !m.flags.Silent {
		fmt.Println(diff)
	}
	if m.flags.Mail {
		err = m.sendMail(diff)
		if err != nil {
			log.Error(err)
		}
	}

	err = store.Save(m.domain)
	if err != nil {
		log.Fatal(err)
	}
}

func (m Monitor) sendMail(diff string) error {
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

func (m Monitor) getDiff(answers []string) (string, error) {
	lo := m.domain.GetLastObservation()
	diff := difflib.UnifiedDiff{
		A:        lo.GetAnswers(),
		B:        answers,
		FromFile: "Last answer " + lo.Time().String(),
		ToFile:   "Current answer " + time.Now().String(),
		Context:  1,
	}
	return difflib.GetUnifiedDiffString(diff)
}

// Observe queries DNS and creates a Record of observed answers
func (m Monitor) Observe() store.Record {
	msg := dns.Msg{}
	msg.SetQuestion(m.domain.Name+".", dns.TypeA)
	dnsClient := dns.Client{}
	r, _, err := dnsClient.Exchange(&msg, m.flags.DNS+":53")
	if err != nil {
		log.Fatal(err)
	}

	answers := []string{}
	for _, a := range r.Answer {
		answers = append(answers, strings.Fields(a.String())[4])
	}

	record := store.CreateRecord(answers)
	m.domain.Observations = append(m.domain.Observations, record)
	return record
}
