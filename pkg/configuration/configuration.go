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
		configFile = readConfig(data)
		for name, check := range configFile.Checks {

			configMap[name] = Config{
				Domains:  check.Names,
				DNS:      flags.DNS,
				Silent:   flags.Silent,
				Interval: flags.Interval,
				Mail:     flags.Mail,
				SMS:      flags.SMS,
			}
		}
	} else {
		configMap["default"] = Config{
			Domains:  flags.Domains,
			DNS:      flags.DNS,
			Silent:   flags.Silent,
			Interval: flags.Interval,
			Mail:     flags.Mail,
			SMS:      flags.SMS,
		}
	}

	return configMap
}
