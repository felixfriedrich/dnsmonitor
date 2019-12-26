package config

import (
	"bytes"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// CreateEnvConfigFromEnv returns a valid config or returns an error
func CreateEnvConfigFromEnv(prefix string, config interface{}) error {
	err := envconfig.Process(prefix, config)

	if err != nil {
		envconfig.Usage(prefix, config)
		buffer := bytes.NewBufferString("")
		envconfig.Usagef(prefix, config, buffer, envconfig.DefaultTableFormat)
		log.Error("failed to create config from env")
		log.Error("found config ", config)
		log.Error(buffer.String())
		return errors.Wrap(err, buffer.String())
	}

	return nil
}
