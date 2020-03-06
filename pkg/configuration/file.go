package configuration

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// ConfigFile represents the yml structure expected in a configuration file
type ConfigFile struct {
	Monitors map[string]Monitor `yaml:"monitors"`
}

// Monitor is part of the yml structure expected in a configuration file
type Monitor struct {
	Domains  []string `yaml:"domains"`
	DNS      string   `yaml:"dns"`
	Interval int      `yaml:"interval"`
	Mail     bool     `yaml:"mail"`
	SMS      bool     `yaml:"sms"`
	Silent   bool     `yaml:"silent"`
	Alerting Alerting `yaml:"alerting"`
}

type Alerting struct {
	Mail MailAlerting `yaml:"mail"`
}

type MailAlerting struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	To       string `yaml:"to"`
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
	for name, ymlFile := range configFile.Monitors {

		config[name] = Check{
			Domains:  ymlFile.Domains,
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
