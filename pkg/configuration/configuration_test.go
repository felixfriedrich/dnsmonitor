package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfigFromFlagsOnly(t *testing.T) {
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

func TestCreateConfig_ConfigFileOverridesDNSFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		DNS:        "8.8.8.8",
	}
	config := Create(flags)["amazon"]
	assert.Equal(t, "8.8.4.4", config.DNS)
}

// This test assures the DNS flag is properly used, if there is nothing in the config file
func TestCreateConfig_DNSFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		DNS:        "8.8.8.8",
	}
	config := Create(flags)["default"]
	assert.Equal(t, "8.8.8.8", config.DNS)
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
