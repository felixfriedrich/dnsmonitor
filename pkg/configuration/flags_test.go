package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainFlag(t *testing.T) {
	d := Domains{"a", "b", "c"}
	assert.Equal(t, Domains{"a", "b", "c"}, d)
	assert.Equal(t, "[a, b, c]", d.String())
}

func TestVendorFlag(t *testing.T) {
	vf := VendorFlag{}
	vf.Set("messagebird")
	assert.Equal(t, MessageBird, vf.Vendor)
	assert.Equal(t, "messagebird", vf.String())
}
