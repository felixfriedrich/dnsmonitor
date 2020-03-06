package configuration

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func parseYml(data []byte) Config {
	var config Config
	err := yaml.UnmarshalStrict(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
