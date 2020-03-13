package configuration

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
)

const (
	// EnvMailPrefix is used to prefix env vars
	EnvMailPrefix = "dnsmonitor_mail"
)

func mergeFlags(config Config, flags Flags) Config {

	if _, exists := config.Monitors[Default]; !exists {
		config.Monitors[Default] = &Monitor{}
	}

	for name, ymlFile := range config.Monitors {
		config.Monitors[name] = &Monitor{
			/*
				The main reason to merge the fields is, that the flag value serves as a default.
				Even if the flag is not specified, the flag itself has a default value.
				But that also means that any neutral value (e.g. false), will be overridden by the value of the flag.
			*/
			Domains:  merge(ymlFile.Domains, flags.Domains).(Domains),
			DNS:      merge(ymlFile.DNS, flags.DNS).(string),
			Interval: merge(ymlFile.Interval, flags.Interval).(int),
			Mail:     merge(ymlFile.Mail, flags.Mail).(bool),
			SMS:      merge(ymlFile.SMS, flags.SMS).(bool),
			Silent:   merge(ymlFile.Silent, flags.Silent).(bool),
			Alerting: merge(ymlFile.Alerting, flags.VendorFlag.Vendor).(Alerting),
		}
	}
	return config
}

func mergeEnvVars(config Config) (Config, error) {
	for name, ymlFile := range config.Monitors {
		if ymlFile.Mail && ymlFile.Alerting.Mail.To == "" {
			mailConfig := config.Monitors[name].Alerting.Mail
			err := envconfig.Process(EnvMailPrefix, &mailConfig)
			if err != nil {
				return config, errors.New("alerting mail config could neither be loaded from config file nor from env vars")
			}
			config.Monitors[name].Alerting.Mail = mailConfig
		}
	}
	return config, nil
}

func merge(value interface{}, defaultValue interface{}) interface{} {
	s, ok := value.(string)
	if ok && s != "" {
		return value
	}

	i, ok := value.(int)
	if ok && i != 0 {
		return value
	}

	b, ok := value.(bool)
	if ok && b != false {
		return value
	}

	domains, ok := value.(Domains)
	if ok && len(domains) > 0 {
		return value
	}

	alerting, ok := value.(Alerting)
	if ok && alerting.SMS.Vendor != None {
		return alerting
	} // using else here instead didn't work. The return statement at the end became unreachable code.
	if ok && alerting.SMS.Vendor == None {
		alerting.SMS.Vendor = defaultValue.(Vendor)
		return alerting
	}

	return defaultValue
}
