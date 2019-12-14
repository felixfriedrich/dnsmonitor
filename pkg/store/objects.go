package store

import (
	"net"
	"sort"
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
