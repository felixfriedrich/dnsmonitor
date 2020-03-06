package envconfig

import (
	"bytes"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Read reads environment variables with the help of the envconfig package
func Read(prefix string, config interface{}) {
	err := envconfig.Process(prefix, config)
	if err != nil {
		usage := bytes.NewBufferString("")
		envconfig.Usagef(prefix, config, usage, envconfig.DefaultTableFormat)
		fmt.Println(usage)

		fmt.Println("Found config: ", config)
		os.Exit(1)
	}
}
