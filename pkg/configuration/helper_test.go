package configuration

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeWithString(t *testing.T) {
	assert.Equal(t, "Hello", merge("Hello", "World"))
	assert.Equal(t, "World", merge("", "World"))
}

func TestMergeWithInt(t *testing.T) {
	assert.Equal(t, 1, merge(0, 1))
	assert.Equal(t, 1, merge(1, 0))
}

func TestMergeWithBool(t *testing.T) {
	assert.True(t, merge(false, true).(bool))
	assert.True(t, merge(true, false).(bool))
}

func TestMergeWithDomainList(t *testing.T) {
	assert.Equal(t, Domains{"www.google.com"}, merge(Domains{"www.google.com"}, Domains{}).(Domains))
	assert.Equal(t, Domains{"www.google.com"}, merge(Domains{}, Domains{"www.google.com"}).(Domains))
}

func TestMergeVendorFlagIntoAlerting(t *testing.T) {
	a := Alerting{}
	v := MessageBird
	assert.Equal(t, MessageBird, merge(a, v).(Alerting).SMS.Vendor)
}

func TestMergeVendorFromConfigFileOverridesFlag(t *testing.T) {
	a := Alerting{
		SMS: SMSAlerting{Vendor: MessageBird},
	}
	v := SMS77
	assert.Equal(t, MessageBird, merge(a, v).(Alerting).SMS.Vendor)
}

func TestMergeEnvVars_NoMailConfigNoEnvVars(t *testing.T) {
	config := NewConfig()
	monitor := Monitor{
		Mail: true,
		Alerting: Alerting{Mail: MailAlerting{
			To: "",
		}},
	}
	config.Monitors["default"] = &monitor

	_, err := mergeEnvVars(config)
	assert.Error(t, err)
}

func TestMergeEnvVars_NoMailConfigButEnvVars(t *testing.T) {
	config := NewConfig()
	localhost := "localhost"
	port25 := 25
	root := "root"
	password := "1234"
	from := "dnsmonitor@localhost"
	to := "me@localhost"
	os.Setenv("DNSMONITOR_MAIL_HOST", localhost)
	os.Setenv("DNSMONITOR_MAIL_PORT", strconv.Itoa(port25))
	os.Setenv("DNSMONITOR_MAIL_USERNAME", root)
	os.Setenv("DNSMONITOR_MAIL_PASSWORD", password)
	os.Setenv("DNSMONITOR_MAIL_FROM", from)
	os.Setenv("DNSMONITOR_MAIL_TO", to)

	monitor := Monitor{
		Mail:     true,
		Alerting: Alerting{Mail: MailAlerting{}},
	}
	config.Monitors["default"] = &monitor

	config, err := mergeEnvVars(config)
	assert.NoError(t, err)
	assert.Equal(t, localhost, config.Monitors["default"].Alerting.Mail.Host)
	assert.Equal(t, port25, config.Monitors["default"].Alerting.Mail.Port)
	assert.Equal(t, root, config.Monitors["default"].Alerting.Mail.Username)
	assert.Equal(t, password, config.Monitors["default"].Alerting.Mail.Password)
	assert.Equal(t, from, config.Monitors["default"].Alerting.Mail.From)
	assert.Equal(t, to, config.Monitors["default"].Alerting.Mail.To)

	os.Unsetenv("DNSMONITOR_MAIL_HOST")
	os.Unsetenv("DNSMONITOR_MAIL_PORT")
	os.Unsetenv("DNSMONITOR_MAIL_USERNAME")
	os.Unsetenv("DNSMONITOR_MAIL_PASSWORD")
	os.Unsetenv("DNSMONITOR_MAIL_FROM")
	os.Unsetenv("DNSMONITOR_MAIL_TO")
}
