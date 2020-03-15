package configuration

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeEnvVars_NoMailConfigNoEnvVars(t *testing.T) {
	config := NewConfig()
	monitor := Monitor{
		Mail: true,
		Alerting: Alerting{Mail: MailAlerting{
			To: "",
		}},
	}
	config.Monitors[Default] = &monitor

	_, err := mergeEnvVars(config)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DNSMONITOR_MAIL_USERNAME")
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
	config.Monitors[Default] = &monitor

	config, err := mergeEnvVars(config)
	assert.NoError(t, err)
	assert.Equal(t, localhost, config.Monitors[Default].Alerting.Mail.Host)
	assert.Equal(t, port25, config.Monitors[Default].Alerting.Mail.Port)
	assert.Equal(t, root, config.Monitors[Default].Alerting.Mail.Username)
	assert.Equal(t, password, config.Monitors[Default].Alerting.Mail.Password)
	assert.Equal(t, from, config.Monitors[Default].Alerting.Mail.From)
	assert.Equal(t, to, config.Monitors[Default].Alerting.Mail.To)

	os.Unsetenv("DNSMONITOR_MAIL_HOST")
	os.Unsetenv("DNSMONITOR_MAIL_PORT")
	os.Unsetenv("DNSMONITOR_MAIL_USERNAME")
	os.Unsetenv("DNSMONITOR_MAIL_PASSWORD")
	os.Unsetenv("DNSMONITOR_MAIL_FROM")
	os.Unsetenv("DNSMONITOR_MAIL_TO")
}

func TestMergeEnvVars_NoMessageBirdConfigNoEnvVars(t *testing.T) {
	config := NewConfig()
	monitor := Monitor{
		SMS: true,
		Alerting: Alerting{SMS: SMSAlerting{
			Vendor: MessageBird,
			MessageBird: MessageBirdConfig{
				AccessKey: "",
			},
		}},
	}
	config.Monitors[Default] = &monitor

	_, err := mergeEnvVars(config)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DNSMONITOR_MESSAGEBIRD_ACCESSKEY")
}

func TestMergeEnvVars_NoMessageBirdConfigButEnvVars(t *testing.T) {
	config := NewConfig()

	os.Setenv("DNSMONITOR_MESSAGEBIRD_ACCESSKEY", "2xgWcAepPYdUGvhH7t1H")
	os.Setenv("DNSMONITOR_MESSAGEBIRD_SENDER", "+49 30 835646496")
	os.Setenv("DNSMONITOR_MESSAGEBIRD_RECIPIENTS", "+49 30 239768508")

	monitor := Monitor{
		SMS: true,
		Alerting: Alerting{SMS: SMSAlerting{
			Vendor: MessageBird,
			MessageBird: MessageBirdConfig{
				AccessKey: "",
			},
		}},
	}
	config.Monitors[Default] = &monitor

	_, err := mergeEnvVars(config)
	assert.NoError(t, err)

	os.Unsetenv("DNSMONITOR_MESSAGEBIRD_ACCESSKEY")
	os.Unsetenv("DNSMONITOR_MESSAGEBIRD_SENDER")
	os.Unsetenv("DNSMONITOR_MESSAGEBIRD_RECIPIENTS")
}

func TestMergeEnvVars_NoSMS77ConfigNoEnvVars(t *testing.T) {
	config := NewConfig()
	monitor := Monitor{
		SMS: true,
		Alerting: Alerting{SMS: SMSAlerting{
			Vendor: SMS77,
			SMS77: SMS77Config{
				APIKey: "",
			},
		}},
	}
	config.Monitors[Default] = &monitor

	_, err := mergeEnvVars(config)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DNSMONITOR_SMS77_APIKEY")
}

func TestMergeEnvVars_NoSMS77ConfigButEnvVars(t *testing.T) {
	config := NewConfig()

	APIKey := "2xgWcAepPYdUGvhH7t1H"
	sender := "+49 30 835646496"
	recipient := "+49 30 239768508"
	debug := false
	os.Setenv("DNSMONITOR_SMS77_APIKEY", APIKey)
	os.Setenv("DNSMONITOR_SMS77_SENDER", sender)
	os.Setenv("DNSMONITOR_SMS77_RECIPIENT", recipient)
	os.Setenv("DNSMONITOR_SMS77_DEBUG", strconv.FormatBool(debug))

	monitor := Monitor{
		SMS: true,
		Alerting: Alerting{SMS: SMSAlerting{
			Vendor: SMS77,
			SMS77: SMS77Config{
				APIKey: "",
			},
		}},
	}
	config.Monitors[Default] = &monitor

	config, err := mergeEnvVars(config)
	assert.NoError(t, err)

	assert.Equal(t, APIKey, config.Monitors[Default].Alerting.SMS.SMS77.APIKey)
	assert.Equal(t, sender, config.Monitors[Default].Alerting.SMS.SMS77.Sender)
	assert.Equal(t, recipient, config.Monitors[Default].Alerting.SMS.SMS77.Recipient)
	assert.Equal(t, debug, config.Monitors[Default].Alerting.SMS.SMS77.Debug)

	os.Unsetenv("DNSMONITOR_SMS77_APIKEY")
	os.Unsetenv("DNSMONITOR_SMS77_SENDER")
	os.Unsetenv("DNSMONITOR_SMS77_RECIPIENT")
	os.Unsetenv("DNSMONITOR_SMS77_DEBUG")
}

func TestMergeEnvVars_SMSFlagButNoVendor(t *testing.T) {
	config := NewConfig()
	monitor := Monitor{
		SMS: true,
		Alerting: Alerting{SMS: SMSAlerting{
			Vendor: SMS77,
			SMS77: SMS77Config{
				APIKey: "",
			},
		}},
	}
	config.Monitors[Default] = &monitor

	_, err := mergeEnvVars(config)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DNSMONITOR_SMS77_APIKEY")
}
