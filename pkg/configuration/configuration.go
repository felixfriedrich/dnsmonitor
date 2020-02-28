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

// Create takes command line flags and converts them into a generic Config object
func Create(flags Flags) Config {
	domains := flags.Domains
	var configFile ConfigFile
	if flags.ConfigFile != "" {
		data, err := ioutil.ReadFile(flags.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		configFile = readConfig(data)
		domains = configFile.Checks["default"].Names
	}

	return Config{
		Domains:  domains,
		DNS:      flags.DNS,
		Silent:   flags.Silent,
		Interval: flags.Interval,
		Mail:     flags.Mail,
		SMS:      flags.SMS,
	}
}
