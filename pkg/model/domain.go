package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/pmezard/go-difflib/difflib"

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

// CreateDomain creates a domain object initialised with the correct name and an empty list of Observations
func CreateDomain(domain string) *Domain {
	return &Domain{Name: domain, Observations: []Record{}}
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
