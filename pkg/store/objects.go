package store

import (
	"errors"
	"github.com/pmezard/go-difflib/difflib"
	"net"
	"sort"
	"strconv"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Domain is used to store a domain name and a list of IP this domain has/had (before)
type Domain struct {
	Name         string
	Observations []Record
}

func (d *Domain) String() string {
	return d.Name + " [" + strconv.Itoa(len(d.Observations)) + " observations]"
}

// Record is used to store the result of the  DNS query and the point in time it has been observed
type Record struct {
	time   time.Time
	cnames []string
	ips    []string
}

// Equal is needed by cmp.Equal in order to compare two Records
// We we don't care about the time the observation has been made (Record has been created), only DNS answers are considered
func (r Record) Equal(rr Record) bool {
	return cmp.Equal(r.GetAnswers(), rr.GetAnswers())
}

// Time returns the timestamp of Record r
func (r *Record) Time() time.Time {
	return r.time
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

// CreateDomain creates a domain object initialised with the correct name and an empty list of Observations
func CreateDomain(domain string) *Domain {
	return &Domain{Name: domain, Observations: []Record{}}
}

// GetAnswers return the answers like they have been observed
func (r *Record) GetAnswers() []string {
	return append(r.cnames, r.ips...)
}

// GetDiff returns the difference between a Record r and the last observation of that domain
func (d *Domain) GetDiff(r Record) (string, error) {
	lo := d.GetLastObservation()
	diff := difflib.UnifiedDiff{
		A:        lo.GetAnswers(),
		B:        r.GetAnswers(),
		FromFile: "Last answer " + lo.Time().String(),
		ToFile:   "Current answer " + time.Now().String(),
		Context:  1,
	}
	return difflib.GetUnifiedDiffString(diff)
}

// GetLastObservation returns the last observation (Record)
func (d *Domain) GetLastObservation() *Record {
	if len(d.Observations) == 0 {
		return &Record{}
	}
	lastElement := len(d.Observations) - 1
	o := d.Observations[lastElement]
	return &o
}

// LastChangeRecord returns the time the last change in DNS answers has been observed
func (d *Domain) LastChangeRecord() (Record, error) {
	if len(d.Observations) == 0 {
		return Record{}, errors.New("no observations made yet")
	}
	r := d.Observations[len(d.Observations)-1] // Return time of first Record in case there is just one Record
	for i := len(d.Observations) - 2; i >= 1; i-- {
		current := d.Observations[i]
		next := d.Observations[i-1]
		equal := cmp.Equal(current, next)
		if !equal {
			return current, nil
		}
	}
	return r, nil
}
