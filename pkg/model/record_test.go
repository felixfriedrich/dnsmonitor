package model

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

func TestGetAnswers(t *testing.T) {
	// dig +short www.isrctn.com
	a := []string{"star.live.cf.public.springer.com.", "prod.springer2.map.fastlylb.net.", "151.101.0.95", "151.101.64.95", "151.101.128.95", "151.101.192.95"}
	r := CreateRecord(a)

	// As the order might have changed from the original input, only lengths are compared
	assert.Len(t, r.GetAnswers(), len(a))
}

func TestCompareRecords(t *testing.T) {
	r1 := *CreateRecord([]string{"93.184.216.34"})
	r2 := *CreateRecord([]string{"93.184.216.34"})

	assert.True(t, cmp.Equal(r1, r2))
}
