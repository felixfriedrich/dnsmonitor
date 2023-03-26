package dns

import (
	"log"
	"strings"

	lib "github.com/miekg/dns"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// Interface abtracts the upstream library away, mainly for mocking
//
//counterfeiter:generate . Interface
type Interface interface {
	Query(domain string, dnsServer string) ([]string, error)
}

type d struct {
	client lib.Client
}

// New returns an interface to the DNS library
func New() Interface {
	return &d{client: lib.Client{}}
}

func (d *d) Query(domain string, dnsServer string) ([]string, error) {
	msg := lib.Msg{}
	msg.SetQuestion(domain+".", lib.TypeA)

	r, _, err := d.client.Exchange(&msg, dnsServer+":53")
	if err != nil {
		return nil, err
	}
	if r == nil {
		log.Fatal("dns Exchange returned nil value")
	}

	answers := []string{}
	for _, a := range r.Answer {
		answers = append(answers, strings.Fields(a.String())[4])
	}
	return answers, nil
}
