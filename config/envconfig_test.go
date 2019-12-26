package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Test string `required:"true"`
}

func TestLoadTestConfig(t *testing.T) {
	os.Setenv("TEST_TEST", "TEST")
	c := &TestConfig{}
	err := CreateEnvConfigFromEnv("test", c)
	assert.NoError(t, err)
	assert.Equal(t, "TEST", c.Test)
	os.Unsetenv("TEST_TEST")
}

func TestLoadTestConfigError(t *testing.T) {
	c := &TestConfig{}
	err := CreateEnvConfigFromEnv("test", c)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "This application is configured via the environment.")
}
