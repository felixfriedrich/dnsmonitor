package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
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

	if err != nil {
		envconfig.Usage(prefix, &config)
		log.Error("failed to create mail config from env")
		log.Error("Found config ", config)
		os.Exit(1)
	}

	return config
}
