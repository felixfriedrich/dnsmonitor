package alerting

import (
	"dnsmonitor/pkg/alerting/messagebird"
	"dnsmonitor/pkg/alerting/sms77"
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/configuration/envconfig"
	"errors"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// API abstracts the upstream library away, mainly for mocking
//counterfeiter:generate . API
type API interface {
	SendSMS(text string) error
}

// Type of message to send
type Type uint8

const (
	// SMS ...
	SMS Type = 0
)

// New returns a new alerting API using a specific implementation
func New(vendor configuration.Vendor, t Type) (API, error) {
	var alertingAPI API
	if vendor == configuration.MessageBird && t == SMS {
		c := messagebird.Config{}
		prefix := "dnsmonitor_messagebird"
		envconfig.Read(prefix, &c)
		return messagebird.New(c), nil
	}
	if vendor == configuration.SMS77 && t == SMS {
		c := sms77.Config{}
		prefix := "dnsmonitor_sms77"
		envconfig.Read(prefix, &c)
		return sms77.New(c), nil
	}
	return alertingAPI, errors.New("vendor/type combination isn't supported")
}
