package configuration

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

const (
	// Default is the name of the default monitor in config files
	Default = "default"
	// EnvMailPrefix is used to prefix env vars
	EnvMailPrefix = "dnsmonitor_mail"
	// EnvMessageBirdPrefix is used to prefix env vars
	EnvMessageBirdPrefix = "dnsmonitor_messagebird"
	// EnvSMS77Prefix is used to prefix env vars
	EnvSMS77Prefix = "dnsmonitor_sms77"
)

type condition func() bool

// CreateConfig takes command line flags, reads a config file (if specified) and returns a config.
// Flags are merged into the information from the config file.
func CreateConfig(flags Flags) (Config, error) {
	config := NewConfig()

	if flags.ConfigFile != "" {
		data, err := ioutil.ReadFile(flags.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		config = parseYml(data)
	}

	config = mergeFlags(config, flags)
	return mergeEnvVars(config)
}

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
