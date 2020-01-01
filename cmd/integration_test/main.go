package main

import (
	"dnsmonitor/pkg/alerting"
	"fmt"
)

func main() {

	alertingAPI := alerting.New(alerting.MessageBird, alerting.SMS)
	err := alertingAPI.SendSMS("Test")
	if err != nil {
		panic(err)
	}

	fmt.Println("OK")
}
