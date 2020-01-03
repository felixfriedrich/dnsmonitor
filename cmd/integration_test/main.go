package main

import (
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/configuration"
	"flag"
	"fmt"
)

type flags struct {
	VendorFlag configuration.VendorFlag
}

func main() {
	f := flags{}
	flag.Var(&f.VendorFlag, "vendor", "alerting vendor, e.g. 'messagebird'.")
	flag.Parse()

	fmt.Println("Using vendor:", f.VendorFlag.String())

	if f.VendorFlag.Vendor == alerting.MessageBird {
		alertingAPI := alerting.New(alerting.MessageBird, alerting.SMS)
		err := alertingAPI.SendSMS("Test")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("OK")
}
