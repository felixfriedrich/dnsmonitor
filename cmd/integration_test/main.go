package main

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/alerting"
	"flag"
	"fmt"
)

type flags struct {
	VendorFlag config.VendorFlag
}

func main() {
	f := flags{}
	flag.Var(&f.VendorFlag, "vendor", "alerting vendor, e.g. 'messagebird'.")
	flag.Parse()

	fmt.Println("Using vendor:", f.VendorFlag)

	if f.VendorFlag.Vendor == alerting.MessageBird {
		alertingAPI := alerting.New(alerting.MessageBird, alerting.SMS)
		err := alertingAPI.SendSMS("Test")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("OK")
}
