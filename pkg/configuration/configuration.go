package configuration

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// Config represents a configuration (independent from flags or environment variables)
type Config map[string]Check

// Check holds information needed for one check performed on n domains.
type Check struct {
	Domains  Domains
	DNS      string
	Silent   bool
	Interval int
	Mail     bool
	SMS      bool
}

// CreateConfig takes command line flags and converts them into a map of config objects also reading a config file, if specified
func CreateConfig(flags Flags) Config {
	var configFile ConfigFile

	if flags.ConfigFile != "" {
		data, err := ioutil.ReadFile(flags.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		configFile = parseYml(data)
		return mergeFlags(configFile, flags)
	}

	return createConfigFromFlags(flags)
}
