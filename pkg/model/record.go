package model

import (
	"net"
	"sort"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Record is used to store the result of the  DNS query and the point in time it has been observed
type Record struct {
	time   time.Time
	cnames []string
	ips    []string
}

// CreateRecord creates a record with the current time as timestamp
func CreateRecord(answers []string) *Record {
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
	return &Record{time: time.Now(), cnames: cnames, ips: ips}
}

// Equal is needed by cmp.Equal in order to compare two Records
// We we don't care about the time the observation has been made (Record has been created), only DNS answers are considered
func (r Record) Equal(rr Record) bool {
	return cmp.Equal(r.GetAnswers(), rr.GetAnswers())
}

// GetAnswers return the answers like they have been observed
func (r *Record) GetAnswers() []string {
	return append(r.cnames, r.ips...)
}
