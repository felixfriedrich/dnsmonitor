package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeWithString(t *testing.T) {
	assert.Equal(t, "Hello", merge("Hello", "World"))
	assert.Equal(t, "World", merge("", "World"))
}

func TestMergeWithInt(t *testing.T) {
	assert.Equal(t, 1, merge(0, 1))
	assert.Equal(t, 1, merge(1, 0))
}

func TestMergeWithBool(t *testing.T) {
	assert.True(t, merge(false, true).(bool))
	assert.True(t, merge(true, false).(bool))
}

func TestMergeWithDomainList(t *testing.T) {
	assert.Equal(t, Domains{"www.google.com"}, merge(Domains{"www.google.com"}, Domains{}).(Domains))
	assert.Equal(t, Domains{"www.google.com"}, merge(Domains{}, Domains{"www.google.com"}).(Domains))
}
