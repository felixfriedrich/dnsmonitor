package configuration

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// CreateConfig takes command line flags, reads a config file (if specified) and returns a config.
// Flags are merged into the information from the config file.
func CreateConfig(flags Flags) Config {
	config := NewConfig()

	if flags.ConfigFile != "" {
		data, err := ioutil.ReadFile(flags.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		config = parseYml(data)
	}

	return mergeFlags(config, flags)
}
