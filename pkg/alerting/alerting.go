package alerting

import (
	"dnsmonitor/pkg/alerting/messagebird"
	"dnsmonitor/pkg/alerting/sms77"
	"dnsmonitor/pkg/configuration/envconfig"
	"errors"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// API abstracts the upstream library away, mainly for mocking
//counterfeiter:generate . API
type API interface {
	SendSMS(text string) error
}

// Vendor identifies services for alerting
type Vendor uint8

const (
	// None is used as default by the 'flags' package
	None Vendor = 0
	// MessageBird https://www.messagebird.com/en/
	MessageBird Vendor = 1
	// SMS77 https://app.sms77.io
	SMS77 Vendor = 2
)

// Type of message to send
type Type uint8

const (
	// SMS ...
	SMS Type = 0
)

// New returns a new alerting API using a specific implementation
func New(vendor Vendor, t Type) (API, error) {
	var alertingAPI API
	if vendor == MessageBird && t == SMS {
		c := messagebird.Config{}
		prefix := "dnsmonitor_messagebird"
		envconfig.Read(prefix, &c)
		return messagebird.New(c), nil
	}
	if vendor == SMS77 && t == SMS {
		c := sms77.Config{}
		prefix := "dnsmonitor_sms77"
		envconfig.Read(prefix, &c)
		return sms77.New(c), nil
	}
	return alertingAPI, errors.New("vendor/type combination isn't supported")
}
