package configuration

func mergeFlags(config Config, flags Flags) Config {
	for name, ymlFile := range config.Monitors {
		config.Monitors[name] = Monitor{
			Domains:  ymlFile.Domains, // Domains are not merged as the domain field is not merge in yml
			DNS:      merge(ymlFile.DNS, flags.DNS).(string),
			Interval: merge(ymlFile.Interval, flags.Interval).(int),
			Mail:     merge(ymlFile.Mail, flags.Mail).(bool),
			SMS:      merge(ymlFile.SMS, flags.SMS).(bool),
			Silent:   merge(ymlFile.Silent, flags.Silent).(bool),
			Alerting: Alerting{},
		}
	}
	return config
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
	return defaultValue
}
