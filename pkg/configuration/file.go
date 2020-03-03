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
	Names []string `yaml:"names"`
}

func fromYml(data []byte) ConfigFile {
	var config ConfigFile
	err := yaml.UnmarshalStrict(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
