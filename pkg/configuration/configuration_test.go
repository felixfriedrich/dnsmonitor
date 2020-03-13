package configuration

/*
This file only contains tests, which assure the config is read correctly. As in, information from the config file,
command line flags and env variables are processed with the right priority.
*/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfigFromFlagsOnly(t *testing.T) {
	flags := Flags{
		Domains:    Domains{"www.google.com", "google.com"},
		DNS:        "8.8.8.8",
		Silent:     false,
		Interval:   300,
		Mail:       false,
		SMS:        false,
		Version:    false,
		VendorFlag: VendorFlag{None},
	}
	config, err := CreateConfig(flags)

	assert.NoError(t, err)
	assert.Len(t, config.Monitors, 1)
	defaultConfig := config.Monitors[Default]

	assert.Equal(t, flags.Domains, defaultConfig.Domains)
	assert.Equal(t, flags.DNS, defaultConfig.DNS)
	assert.Equal(t, flags.Silent, defaultConfig.Silent)
	assert.Equal(t, flags.Interval, defaultConfig.Interval)
	assert.Equal(t, flags.Mail, defaultConfig.Mail)
}

func TestCreateConfig_DomainsFromFileOverridesFlags(t *testing.T) {
	examplecom := "example.com"
	flags := Flags{
		Domains:    []string{examplecom},
		ConfigFile: "../../test/config.yml",
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors[Default]
	assert.Contains(t, config.Domains, "google.com")
	assert.Contains(t, config.Domains, "www.google.com")
	assert.NotContains(t, config.Domains, examplecom)
}

func TestCreateConfig_ConfigFileOverridesDNSFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		DNS:        "8.8.8.8",
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors["amazon"]
	assert.Equal(t, "8.8.4.4", config.DNS)
}

// This test assures the DNS flag is properly used, if there is nothing in the config file
func TestCreateConfig_DNSFlag(t *testing.T) {
	dnsServer := "8.8.8.8"
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		DNS:        dnsServer,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors[Default]
	assert.Equal(t, dnsServer, config.DNS)
}

func TestCreateConfig_ConfigFileOverridesIntervalFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Interval:   300,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors["amazon"]
	assert.Equal(t, 5, config.Interval)
}

// This test assures the interval flag is properly used, if there is nothing in the config file
func TestCreateConfig_IntervalFlag(t *testing.T) {
	interval := 300
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Interval:   interval,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors[Default]
	assert.Equal(t, interval, config.Interval)
}

func TestCreateConfig_ConfigFileOverridesMailFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Mail:       false,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors["amazon"]
	assert.Equal(t, true, config.Mail)
}

// This test assures the mail flag is properly used, if there is nothing in the config file
func TestCreateConfig_MailFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Mail:       false,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors[Default]
	assert.Equal(t, false, config.Mail)
}

func TestCreateConfig_ConfigFileOverridesSMSFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		SMS:        false,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors["amazon"]
	assert.Equal(t, true, config.SMS)
}

// This test assures the sms flag is properly used, if there is nothing in the config file
func TestCreateConfig_SMSFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		SMS:        false,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors[Default]
	assert.Equal(t, false, config.SMS)
}

func TestCreateConfig_ConfigFileOverridesSilentFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Silent:     false,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors["amazon"]
	assert.Equal(t, true, config.Silent)
}

// This test assures the silent flag is properly used, if there is nothing in the config file
func TestCreateConfig_SilentFlag(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		SMS:        false,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors[Default]
	assert.Equal(t, false, config.Silent)
}

func TestNewConfig_MailConfig(t *testing.T) {
	flags := Flags{
		ConfigFile: "../../test/config.yml",
		Mail:       true,
	}
	c, _ := CreateConfig(flags)
	config := c.Monitors["amazon"]
	assert.NotEqual(t, "", config.Alerting.Mail.To)
}
