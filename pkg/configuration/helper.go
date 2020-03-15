package configuration

import (
	"bytes"
	"errors"

	"github.com/kelseyhightower/envconfig"
)

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
