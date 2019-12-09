package store

import (
	"net"
	"sort"
	"time"
)

// Domain is used to store a domain name and a list of IP this domain has/had (before)
type Domain struct {
	name         string
	observations []Record
}

// Record is used to store the result of the  DNS query and the point in time it has been observed
type Record struct {
	time   time.Time
	cnames []string
	ips    []string
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
