package store

import (
	"testing"

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
