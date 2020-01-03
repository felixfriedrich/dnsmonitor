package configuration

import (
	"dnsmonitor/pkg/alerting"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainFlag(t *testing.T) {
	d := DomainFlag{"a", "b", "c"}
	assert.Equal(t, DomainFlag{"a", "b", "c"}, d)
	assert.Equal(t, "[a, b, c]", d.String())
}

func TestVendorFlag_Set(t *testing.T) {
	vf := VendorFlag{}
	vf.Set("messagebird")
	assert.Equal(t, alerting.MessageBird, vf.Vendor)
}