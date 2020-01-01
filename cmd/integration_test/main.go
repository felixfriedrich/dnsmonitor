package main

import "dnsmonitor/pkg/alerting"

func main() {
	alertingAPI := alerting.New(alerting.MessageBird, alerting.SMS)
	alertingAPI.SendSMS("Test")
}
