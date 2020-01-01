package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// ReadEnvConfig reads environment variables with the help of the enconfig package
func ReadEnvConfig(prefix string, config interface{}) {
	err := envconfig.Process(prefix, config)
	if err != nil {
		usage := bytes.NewBufferString("")
		envconfig.Usagef(prefix, config, usage, envconfig.DefaultTableFormat)
		fmt.Println(usage)

		fmt.Println("Found config: ", config)
		os.Exit(1)
	}
}
