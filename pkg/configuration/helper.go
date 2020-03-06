package configuration

func mergeFlags(config Config, flags Flags) Config {
	for name, ymlFile := range config.Monitors {
		config.Monitors[name] = Monitor{
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
