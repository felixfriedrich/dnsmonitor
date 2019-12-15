package store

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestCreateRecordEnsureARecordsAreSorted(t *testing.T) {
	// dig +short isrctn.com
	a := []string{"96.45.82.109", "96.45.83.39", "96.45.83.224", "96.45.82.248"}
	ra := CreateRecord(a)

	b := []string{"96.45.82.248", "96.45.82.109", "96.45.83.39", "96.45.83.224"}
	rb := CreateRecord(b)

	assert.Equal(t, ra.ips, rb.ips)
}
func TestCreateRecordWithCNAMEsAndARecords(t *testing.T) {
	// dig +short www.isrctn.com
	a := []string{"star.live.cf.public.springer.com.", "prod.springer2.map.fastlylb.net.", "151.101.0.95", "151.101.64.95", "151.101.128.95", "151.101.192.95"}
	r := CreateRecord(a)
	assert.Len(t, r.cnames, 2)
	assert.Len(t, r.ips, 4)
}

func TestCreateDomain(t *testing.T) {
	d := CreateDomain("www.google.com")
	assert.Equal(t, "www.google.com", d.Name)
	assert.Equal(t, 0, len(d.Observations))
}

func TestGetAnswers(t *testing.T) {
	// dig +short www.isrctn.com
	a := []string{"star.live.cf.public.springer.com.", "prod.springer2.map.fastlylb.net.", "151.101.0.95", "151.101.64.95", "151.101.128.95", "151.101.192.95"}
	r := CreateRecord(a)

	// As the order might have changed from the original input, only lengths are compared
	assert.Len(t, r.GetAnswers(), len(a))
}

func TestGetLastObservation(t *testing.T) {
	d := CreateDomain("example.com")
	d.Observations = append(d.Observations, *CreateRecord([]string{"93.184.216.34"}))
	d.Observations = append(d.Observations, *CreateRecord([]string{"www.example.com", "93.184.216.34"}))

	o := d.GetLastObservation()
	assert.Equal(t, o.GetAnswers(), []string{"www.example.com", "93.184.216.34"}, o)
}

func TestGetLastAnswerDoesntExist(t *testing.T) {
	d := CreateDomain("example.com")
	o := d.GetLastObservation()
	assert.Len(t, o.GetAnswers(), 0)
}

func TestCompareRecords(t *testing.T) {
	r1 := *CreateRecord([]string{"93.184.216.34"})
	r2 := *CreateRecord([]string{"93.184.216.34"})

	assert.True(t, cmp.Equal(r1, r2))
}

// If there have been several changes been observed already, the time of the first Record after the last change should
// be returned
func TestLastChangeTimeDefault(t *testing.T) {
	d := CreateDomain("example.com")
	r1 := *CreateRecord([]string{"93.184.216.31"})
	r2 := *CreateRecord([]string{"93.184.216.34"})
	r3 := *CreateRecord([]string{"93.184.216.34"})
	r4 := *CreateRecord([]string{"93.184.216.35"})
	d.Observations = append(d.Observations, r1, r2, r3, r4)
	lastChange, err := d.LastChangeTime()
	assert.NoError(t, err)
	assert.Equal(t, r2.time, lastChange)
}

// If there has only one Observation made yet, the time of that should be returned
func TestLastChangeTimeFirstEntry(t *testing.T) {
	d := CreateDomain("example.com")
	r1 := *CreateRecord([]string{"93.184.216.31"})
	d.Observations = append(d.Observations, r1)
	lastChange, err := d.LastChangeTime()
	assert.NoError(t, err)
	assert.Equal(t, r1.time, lastChange)
}

// If this is the fist change to be observed we want the time of the last Observation to be returned
func TestLastChangeTimeFirstChange(t *testing.T) {
	d := CreateDomain("example.com")
	r1 := *CreateRecord([]string{"93.184.216.33"})
	r2 := *CreateRecord([]string{"93.184.216.33"})
	r3 := *CreateRecord([]string{"93.184.216.33"})
	r4 := *CreateRecord([]string{"93.184.216.34"})
	d.Observations = append(d.Observations, r1, r2, r3, r4)
	lastChange, err := d.LastChangeTime()
	assert.NoError(t, err)
	assert.Equal(t, r4.time, lastChange)
}

func TestLastChangeTimeOnlyChanges(t *testing.T) {
	d := CreateDomain("example.com")
	r1 := *CreateRecord([]string{"93.184.216.31"})
	r2 := *CreateRecord([]string{"93.184.216.32"})
	r3 := *CreateRecord([]string{"93.184.216.33"})
	r4 := *CreateRecord([]string{"93.184.216.34"})
	d.Observations = append(d.Observations, r1, r2, r3, r4)
	lastChange, err := d.LastChangeTime()
	assert.NoError(t, err)
	assert.Equal(t, r3.time, lastChange)
}

func TestLastChangeTimeNoObservation(t *testing.T) {
	d := CreateDomain("example.com")
	_, err := d.LastChangeTime()
	assert.Error(t, err)
}
