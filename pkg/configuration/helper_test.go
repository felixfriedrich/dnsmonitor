package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionalWithString(t *testing.T) {
	assert.Equal(t, "Hello", optional("Hello", "World"))
	assert.Equal(t, "World", optional("", "World"))
}

func TestOptionalWithInt(t *testing.T) {
	assert.Equal(t, 1, optional(0, 1))
	assert.Equal(t, 1, optional(1, 0))
}

func TestOptionalWithBool(t *testing.T) {
	assert.True(t, optional(false, true).(bool))
	assert.True(t, optional(true, false).(bool))
}

