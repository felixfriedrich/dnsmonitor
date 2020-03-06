package configuration

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// Check holds information needed for one check performed on n domains.
type Check struct {
	Domains  Domains
	DNS      string
	Silent   bool
	Interval int
	Mail     bool
	SMS      bool
}

// Create takes command line flags and converts them into a map of config objects also reading a config file, if specified
func Create(flags Flags) map[string]Check {
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
