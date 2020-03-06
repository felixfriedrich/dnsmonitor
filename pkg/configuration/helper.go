package configuration

func mergeFlags(config Config, flags Flags) Config {
	for name, ymlFile := range config.Monitors {
		config.Monitors[name] = Monitor{
			Domains:  ymlFile.Domains, // Domains are not merged as the domain field is not optional in yml
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


