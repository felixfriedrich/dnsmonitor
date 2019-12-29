package config

import (
	"bytes"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// HandleEnvConfigError handles errors returned by the envconfig package
func HandleEnvConfigError(err error, config interface{}, prefix string) {
	if err != nil {
		envconfig.Usage(prefix, config)
		buffer := bytes.NewBufferString("")
		envconfig.Usagef(prefix, config, buffer, envconfig.DefaultTableFormat)
		log.Error("failed to create config from env")
		log.Error("found config ", config)
		log.Error(buffer.String())
		log.Fatal(errors.Wrap(err, buffer.String()))
	}
}
