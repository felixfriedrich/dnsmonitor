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

func mergeFlags(config Config, flags Flags) Config {
	for name, ymlFile := range config.Monitors {
		config.Monitors[name] = Monitor{
			Domains:  ymlFile.Domains,
			DNS:      optional(ymlFile.DNS, flags.DNS).(string),
			Interval: optional(ymlFile.Interval, flags.Interval).(int),
			Mail:     optional(ymlFile.Mail, flags.Mail).(bool),
			SMS:      optional(ymlFile.SMS, flags.SMS).(bool),
			Silent:   optional(ymlFile.Silent, flags.Silent).(bool),
			Alerting: Alerting{},
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
