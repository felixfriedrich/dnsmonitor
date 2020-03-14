package configuration

import (
	"bytes"
	"errors"

	"github.com/kelseyhightower/envconfig"
)

const (
	// EnvMailPrefix is used to prefix env vars
	EnvMailPrefix = "dnsmonitor_mail"
	// EnvMessageBirdPrefix is used to prefix env vars
	EnvMessageBirdPrefix = "dnsmonitor_messagebird"
	// EnvSMS77Prefix is used to prefix env vars
	EnvSMS77Prefix = "dnsmonitor_sms77"
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

type condition func() bool

func mergeEnvVars(config Config) (Config, error) {
	var err error

	for name, ymlFile := range config.Monitors {

		// Mail config
		err = readEnv(
			func() bool {
				return ymlFile.Mail && ymlFile.Alerting.Mail.To == ""
			},
			EnvMailPrefix,
			&config.Monitors[name].Alerting.Mail,
			err,
		)

		// Messagebird config
		err = readEnv(
			func() bool {
				return ymlFile.SMS && ymlFile.Alerting.SMS.Vendor == MessageBird && ymlFile.Alerting.SMS.MessageBird.AccessKey == ""
			},
			EnvMessageBirdPrefix,
			&config.Monitors[name].Alerting.SMS.MessageBird,
			err,
		)

		// SMS77
		err = readEnv(
			func() bool {
				return ymlFile.SMS && ymlFile.Alerting.SMS.Vendor == SMS77 && ymlFile.Alerting.SMS.SMS77.APIKey == ""
			},
			EnvSMS77Prefix,
			&config.Monitors[name].Alerting.SMS.SMS77,
			err,
		)

		/*
			Any call of readEnv will either set the error or pass it through. This way evaluating it once at the end of
			this function is enough.
		*/
		if err != nil {
			return config, err
		}
	}
	return config, nil
}

func readEnv(condition condition, prefix string, configStruct interface{}, e error) error {
	if condition() {
		err := envconfig.Process(prefix, configStruct)
		if err != nil {
			usage := bytes.NewBufferString("")
			envconfig.Usagef(prefix, configStruct, usage, envconfig.DefaultTableFormat)
			message := "config could neither be loaded from config file nor from env vars"
			message = message + "\n" + usage.String()
			return errors.New(message)
		}
	}

	return e
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
