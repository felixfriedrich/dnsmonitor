package configuration

import (
	"io/ioutil"
	"log"
)

// Config holds configuration.
type Config struct {
	Domains  Domains
	DNS      string
	Silent   bool
	Interval int
	Mail     bool
	SMS      bool
}

// Create takes command line flags and converts them into a map of config objects also reading a config file, if specified
func Create(flags Flags) map[string]Config {
	configMap := make(map[string]Config)
	var configFile ConfigFile

	if flags.ConfigFile != "" {
		data, err := ioutil.ReadFile(flags.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		configFile = parseYml(data)
		configMap = fromConfigFile(configFile, flags)
	} else {
		configMap = fromFlags(flags)
	}

	return configMap
}
