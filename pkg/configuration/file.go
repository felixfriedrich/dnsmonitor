package configuration

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type monitors map[string]Monitor

// Config represents the yml structure expected in a configuration file
type Config struct {
	Monitors monitors `yaml:"monitors"`
}

// NewConfig created a new Config object and initialised the list of monitors
func NewConfig() Config {
	monitors := make(monitors)
	return Config{Monitors: monitors}
}

// Monitor is part of the yml structure expected in a configuration file
type Monitor struct {
	Domains  Domains  `yaml:"domains"`
	DNS      string   `yaml:"dns"`
	Interval int      `yaml:"interval"`
	Mail     bool     `yaml:"mail"`
	SMS      bool     `yaml:"sms"`
	Silent   bool     `yaml:"silent"`
	Alerting Alerting `yaml:"alerting"`
}

// Alerting hold information for alerting
type Alerting struct {
	Mail MailAlerting `yaml:"mail"`
}

// MailAlerting holds information for alerting via mail
type MailAlerting struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	To       string `yaml:"to"`
}

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
