package config

import (
	"github.com/kelseyhightower/envconfig"
)

var (
	prefix = "dnsmonitor_mail"
)

// MailConfig defines environment variables to configure the mail server to use
type MailConfig struct {
	Host     string `default:"127.0.0.1"`
	Port     int    `default:"25"`
	Username string `required:"true"`
	Password string `required:"true"`
	From     string `required:"true"`
	To       string `required:"true"`
}

// CreateMailConfigFromEnvOrDie returns a valid config or panics
func CreateMailConfigFromEnvOrDie() MailConfig {
	config := MailConfig{}
	err := envconfig.Process(prefix, &config)
	HandleEnvConfigError(err, config)
	return config
}
