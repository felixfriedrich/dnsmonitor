package model

import (
	"errors"
	"strconv"

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

// GetDiff returns the difference between the last observation of that domain and the last record that changed
func (d *Domain) GetDiff() (string, error) {
	lo, err := d.LastChangeRecord()
	if err != nil {
		return err.Error(), nil // using error message as diff for now :-)
	}
	diff := cmp.Diff(lo.GetAnswers(), d.LastObservation().GetAnswers())
	return diff, nil
}

// LastObservation returns the last observation (Record)
func (d *Domain) LastObservation() *Record {
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
	if len(d.Observations) == 1 {
		return d.Observations[len(d.Observations)-1], nil
	}

	for i := len(d.Observations) - 2; i >= 1; i-- {
		current := d.Observations[i]
		next := d.Observations[i-1]
		equal := cmp.Equal(current, next)
		if !equal {
			return current, nil
		}
	}

	// In case the loop doesn't find any change, it must be the first change. So the first record is returned.
	return d.Observations[0], nil
}
