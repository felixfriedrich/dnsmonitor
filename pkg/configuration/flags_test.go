package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainStringConversation(t *testing.T) {
	d := domains{"a", "b", "c"}
	s := d.String()
	assert.Equal(t, "[a, b, c]", s)
}
