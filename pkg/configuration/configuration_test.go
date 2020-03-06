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
	config := Create(flags)

	assert.Len(t, config, 1)
	defaultConfig := config["default"]

	assert.Equal(t, flags.Domains, defaultConfig.Domains)
	assert.Equal(t, flags.DNS, defaultConfig.DNS)
	assert.Equal(t, flags.Silent, defaultConfig.Silent)
	assert.Equal(t, flags.Interval, defaultConfig.Interval)
	assert.Equal(t, flags.Mail, defaultConfig.Mail)
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
	config := Create(flags)
	assert.Len(t, config, 2)
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

func TestCreateConfig_ConfigFileOverridesIntervalFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Interval:   300,
	}
	config := Create(flags)["amazon"]
	assert.Equal(t, 5, config.Interval)
}

// This test assures the interval flag is properly used, if there is nothing in the config file
func TestCreateConfig_IntervalFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Interval:   300,
	}
	config := Create(flags)["default"]
	assert.Equal(t, 300, config.Interval)
}

func TestCreateConfig_ConfigFileOverridesMailFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Mail:       false,
	}
	config := Create(flags)["amazon"]
	assert.Equal(t, true, config.Mail)
}

// This test assures the mail flag is properly used, if there is nothing in the config file
func TestCreateConfig_MailFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Mail:       false,
	}
	config := Create(flags)["default"]
	assert.Equal(t, false, config.Mail)
}

func TestCreateConfig_ConfigFileOverridesSMSFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		SMS:        false,
	}
	config := Create(flags)["amazon"]
	assert.Equal(t, true, config.SMS)
}

// This test assures the sms flag is properly used, if there is nothing in the config file
func TestCreateConfig_SMSFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		SMS:        false,
	}
	config := Create(flags)["default"]
	assert.Equal(t, false, config.SMS)
}

func TestCreateConfig_ConfigFileOverridesSilentFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Silent:     false,
	}
	config := Create(flags)["amazon"]
	assert.Equal(t, true, config.Silent)
}

// This test assures the silent flag is properly used, if there is nothing in the config file
func TestCreateConfig_SilentFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		SMS:        false,
	}
	config := Create(flags)["default"]
	assert.Equal(t, false, config.Silent)
}
