package configuration

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	data, err := ioutil.ReadFile("../../test/config.yml")
	assert.NoError(t, err)
	c := fromYml(data)
	assert.NotNil(t, c)
	assert.True(t, len(c.Checks) > 1)
	amazon := c.Checks["amazon"]
	assert.Len(t, amazon.Names, 2)
	assert.Contains(t, amazon.Names, "aws.amazon.com")
	assert.Contains(t, amazon.Names, "www.amazon.com")
}
