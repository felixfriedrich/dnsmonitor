package configuration

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	data, err := ioutil.ReadFile("../../test/config.yml")
	assert.NoError(t, err)
	c := parseYml(data)
	assert.NotNil(t, c)
	assert.True(t, len(c.Checks) > 1)
	amazon := c.Checks["amazon"]
	assert.Len(t, amazon.Domains, 2)
	assert.Contains(t, amazon.Domains, "aws.amazon.com")
	assert.Contains(t, amazon.Domains, "www.amazon.com")
}

func TestReadAlertingMailConfig(t *testing.T) {
	data, err := ioutil.ReadFile("../../test/config.yml")
	assert.NoError(t, err)
	c := parseYml(data)
	amazon := c.Checks["amazon"]
	assert.NotNil(t, amazon.Alerting.Mail)
}

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
