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
	assert.True(t, len(c.Monitors) > 1)
	amazon := c.Monitors["amazon"]
	assert.Len(t, amazon.Domains, 2)
	assert.Contains(t, amazon.Domains, "aws.amazon.com")
	assert.Contains(t, amazon.Domains, "www.amazon.com")
}

func TestReadAlertingMailConfig(t *testing.T) {
	data, err := ioutil.ReadFile("../../test/config.yml")
	assert.NoError(t, err)
	c := parseYml(data)
	amazon := c.Monitors["amazon"]
	assert.NotNil(t, amazon.Alerting.Mail)
}

