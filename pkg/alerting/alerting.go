package alerting

import (
	"dnsmonitor/pkg/alerting/messagebird"
	"dnsmonitor/pkg/alerting/sms77"
	"dnsmonitor/pkg/configuration"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// API abstracts the upstream library away, mainly for mocking
//counterfeiter:generate . API
type API interface {
	SendSMS(text string) error
}

// New returns a new alerting API using a specific implementation
func New(config *configuration.Monitor) API {
	var alertingAPI API

	if config.SMS && config.Alerting.SMS.Vendor == configuration.MessageBird {
		return messagebird.New(config.Alerting.SMS.MessageBird)
	}

	if config.SMS && config.Alerting.SMS.Vendor == configuration.SMS77 {
		return sms77.New(config.Alerting.SMS.SMS77)
	}

	return alertingAPI
}
