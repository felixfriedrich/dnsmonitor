package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This will not be a one to one mapping in the future, I guess.
func TestCreateConfigFromFlags(t *testing.T) {
	flags := Flags{
		Domains:  []string{"www.google.com", "google.com"},
		DNS:      "8.8.8.8",
		Silent:   false,
		Interval: 300,
		Mail:     false,
		Version:  false,
	}
	configMap := Create(flags)

	assert.Len(t, configMap, 1)
	config := configMap["default"]

	assert.Equal(t, flags.Domains, config.Domains)
	assert.Equal(t, flags.DNS, config.DNS)
	assert.Equal(t, flags.Silent, config.Silent)
	assert.Equal(t, flags.Interval, config.Interval)
	assert.Equal(t, flags.Mail, config.Mail)
}

func TestCreateConfig_DomainsFromFileOverridesFlags(t *testing.T) {
	flags := Flags{
		Domains:    []string{"example.com"},
		ConfigFile: "../../test/config.yml",
	}
	config := Create(flags)["default"]
	assert.Contains(t, config.Domains, "google.com")
	assert.Contains(t, config.Domains, "www.google.com")
	assert.NotContains(t, config.Domains, "example.com")
}

func TestCreateConfig_AllChecksArePresent(t *testing.T) {
	flags := Flags{
		Domains:    []string{"example.com"},
		DNS:        "8.8.8.8",
		Silent:     false,
		Interval:   300,
		Mail:       false,
		Version:    false,
		ConfigFile: "../../test/config.yml",
	}
	configMap := Create(flags)
	assert.Len(t, configMap, 2)
}
