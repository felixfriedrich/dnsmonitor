package store

import (
	"github.com/miekg/dns"
	"log"
	"net"
	"sort"
	"strings"
	"time"
)

// Domain is used to store a domain name and a list of IP this domain has/had (before)
type Domain struct {
	Name         string
	Observations []Record
}

// Record is used to store the result of the  DNS query and the point in time it has been observed
type Record struct {
	time   time.Time
	cnames []string
	ips    []string
}

// Time returns the timestamp of Record r
func (r Record) Time() time.Time {
	return r.time
}

// CreateRecord creates a record with the current time as timestamp
func CreateRecord(answers []string) Record {
	cnames := []string{}
	ips := []string{}
	for _, a := range answers {
		if ip := net.ParseIP(a); ip == nil {
			cnames = append(cnames, a)
		} else {
			ips = append(ips, a)
		}
	}
	sort.Strings(ips)
	return Record{time: time.Now(), cnames: cnames, ips: ips}
}

// CreateDomain creates a domain object initialised with the correct namd and an empty list of Observations
func CreateDomain(domain string) Domain {
	return Domain{Name: domain, Observations: []Record{}}
}

// GetAnswers return the answers like they have been observed
func (r Record) GetAnswers() []string {
	return append(r.cnames, r.ips...)
}

// GetLastObservation returns the last observation (Record)
func (d Domain) GetLastObservation() Record {
	if len(d.Observations) == 0 {
		return Record{}
	}
	lastElement := len(d.Observations) - 1
	o := d.Observations[lastElement]
	return o
}

// Observe queries DNS and creates a Record of observed answers
func (d Domain) Observe() Record {
	m := dns.Msg{}
	m.SetQuestion(d.Name+".", dns.TypeA)
	dnsClient := dns.Client{}
	r, _, err := dnsClient.Exchange(&m, "8.8.8.8:53")
	if err != nil {
		log.Fatal(err)
	}

	answers := []string{}
	for _, a := range r.Answer {
		answers = append(answers, strings.Fields(a.String())[4])
	}

	record := CreateRecord(answers)
	d.Observations = append(d.Observations, record)
	return record
}
