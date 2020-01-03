package configuration

import (
	"dnsmonitor/pkg/alerting"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainStringConversation(t *testing.T) {
	d := DomainFlag{"a", "b", "c"}
	s := d.String()
	assert.Equal(t, "[a, b, c]", s)
}

func TestVendorFlag_Set(t *testing.T) {
	vf := VendorFlag{}
	vf.Set("messagebird")
	assert.Equal(t, alerting.MessageBird, vf.Vendor)
}