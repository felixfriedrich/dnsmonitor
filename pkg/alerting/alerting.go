package alerting

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/alerting/messagebird"

	"github.com/kelseyhightower/envconfig"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// API abtracts the upstream library away, mainly for mocking
//counterfeiter:generate . API
type API interface {
	SendSMS(text string) error
}

// Vendor identifies services for alerting
type Vendor uint8

const (
	// MessageBird https://www.messagebird.com/en/
	MessageBird Vendor = 0
)

// Type of message to send
type Type uint8

const (
	// SMS ...
	SMS Type = 0
)

// New returns a new alerting API using a specific implementation
func New(vendor Vendor, t Type) API {
	var alertingAPI API
	if vendor == MessageBird && t == SMS {
		c := messagebird.Config{}
		prefix := "dnsmonitor_messagebird"
		err := envconfig.Process(prefix, &c)
		config.HandleEnvConfigError(err, c, prefix)
		alertingAPI = messagebird.New(c)
	}
	return alertingAPI
}
