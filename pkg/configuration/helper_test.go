package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionalWithString(t *testing.T) {
	assert.Equal(t, "Hello", merge("Hello", "World"))
	assert.Equal(t, "World", merge("", "World"))
}

func TestOptionalWithInt(t *testing.T) {
	assert.Equal(t, 1, merge(0, 1))
	assert.Equal(t, 1, merge(1, 0))
}

func TestOptionalWithBool(t *testing.T) {
	assert.True(t, merge(false, true).(bool))
	assert.True(t, merge(true, false).(bool))
}

