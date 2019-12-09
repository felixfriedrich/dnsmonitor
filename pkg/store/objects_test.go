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
