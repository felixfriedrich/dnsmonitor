package configuration

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// ConfigFile represents the yml structure expected in a configuration file
type ConfigFile struct {
	Checks Checks `yaml:"checks"`
}

// Checks is part of the yml structure expected in a configuration file
type Checks map[string]Entry

// Entry is part of the yml structure expected in a configuration file
type Entry struct {
	Names    []string `yaml:"names"`
	DNS      string   `yaml:"dns"`
	Interval int      `yaml:"interval"`
	Mail     bool     `yaml:"mail"`
	SMS      bool     `yaml:"sms"`
	Silent   bool     `yaml:"silent"`
}

func parseYml(data []byte) ConfigFile {
	var config ConfigFile
	err := yaml.UnmarshalStrict(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func mergeFlags(configFile ConfigFile, flags Flags) Config {
	config := make(Config)
	for name, ymlFile := range configFile.Checks {

		config[name] = Check{
			Domains:  ymlFile.Names,
			DNS:      optional(ymlFile.DNS, flags.DNS).(string),
			Silent:   optional(ymlFile.Silent, flags.Silent).(bool),
			Interval: optional(ymlFile.Interval, flags.Interval).(int),
			Mail:     optional(ymlFile.Mail, flags.Mail).(bool),
			SMS:      optional(ymlFile.SMS, flags.SMS).(bool),
		}
	}
	return config
}

func optional(optional interface{}, fallback interface{}) interface{} {
	s, ok := optional.(string)
	if ok && s != "" {
		return optional
	}
	i, ok := optional.(int)
	if ok && i != 0 {
		return optional
	}
	b, ok := optional.(bool)
	if ok && b != false {
		return optional
	}
	return fallback
}
