package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDomain(t *testing.T) {
	d := CreateDomain("www.google.com")
	assert.Equal(t, "www.google.com", d.Name)
	assert.Equal(t, 0, len(d.Observations))
}

// If there have been several changes been observed already, the time of the first Record after the last change should
// be returned
func TestDomain_LastChangeRecordDefault(t *testing.T) {
	d := CreateDomain("example.com")
	r1 := *CreateRecord([]string{"93.184.216.31"})
	r2 := *CreateRecord([]string{"93.184.216.34"})
	r3 := *CreateRecord([]string{"93.184.216.34"})
	r4 := *CreateRecord([]string{"93.184.216.35"})
	d.Observations = append(d.Observations, r1, r2, r3, r4)
	lastChange, err := d.LastChangeRecord()
	assert.NoError(t, err)
	assert.Equal(t, r2, lastChange)
}

// If there has only one Observation made yet, the time of that should be returned
func TestDomain_LastChangeRecordFirstEntry(t *testing.T) {
	d := CreateDomain("example.com")
	r1 := *CreateRecord([]string{"93.184.216.31"})
	d.Observations = append(d.Observations, r1)
	lastChange, err := d.LastChangeRecord()
	assert.NoError(t, err)
	assert.Equal(t, r1, lastChange)
}

// If this is the fist change to be observed we want the time of the last Observation to be returned
func TestDomain_LastChangeRecordFirstChange(t *testing.T) {
	d := CreateDomain("example.com")
	r1 := *CreateRecord([]string{"93.184.216.33"})
	r2 := *CreateRecord([]string{"93.184.216.33"})
	r3 := *CreateRecord([]string{"93.184.216.33"})
	r4 := *CreateRecord([]string{"93.184.216.34"})
	d.Observations = append(d.Observations, r1, r2, r3, r4)
	lastChange, err := d.LastChangeRecord()
	assert.NoError(t, err)
	assert.Equal(t, r4, lastChange)
}

func TestDomain_LastChangeRecordOnlyChanges(t *testing.T) {
	d := CreateDomain("example.com")
	r1 := *CreateRecord([]string{"93.184.216.31"})
	r2 := *CreateRecord([]string{"93.184.216.32"})
	r3 := *CreateRecord([]string{"93.184.216.33"})
	r4 := *CreateRecord([]string{"93.184.216.34"})
	d.Observations = append(d.Observations, r1, r2, r3, r4)
	lastChange, err := d.LastChangeRecord()
	assert.NoError(t, err)
	assert.Equal(t, r3, lastChange)
}

func TestDomain_LastChangeRecordNoObservation(t *testing.T) {
	d := CreateDomain("example.com")
	_, err := d.LastChangeRecord()
	assert.Error(t, err)
}

func TestDomain_GetLastObservation(t *testing.T) {
	d := CreateDomain("example.com")
	d.Observations = append(d.Observations, *CreateRecord([]string{"93.184.216.34"}))
	d.Observations = append(d.Observations, *CreateRecord([]string{"www.example.com", "93.184.216.34"}))

	o := d.GetLastObservation()
	assert.Equal(t, o.GetAnswers(), []string{"www.example.com", "93.184.216.34"}, o)
}

func TestDomain_GetLastObservationNoRecords(t *testing.T) {
	d := CreateDomain("example.com")
	o := d.GetLastObservation()
	assert.Len(t, o.GetAnswers(), 0)
}
