package main

import (
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/configuration"
	"flag"
	"fmt"
)

type flags struct {
	VendorFlag configuration.VendorFlag
	Mail       bool
}

func main() {
	f := flags{}
	flag.Var(&f.VendorFlag, "vendor", "alerting vendor, e.g. 'messagebird'.")
	flag.BoolVar(&f.Mail, "mail", false, "Test mail alerting")
	flag.Parse()


	if f.VendorFlag.Vendor == alerting.MessageBird {
		fmt.Println("Using SMS vendor:", f.VendorFlag.String())
		alertingAPI := alerting.New(alerting.MessageBird, alerting.SMS)
		err := alertingAPI.SendSMS("Test")
		if err != nil {
			panic(err)
		}
	}

	if f.Mail {
		mail := alerting.NewMail()
		fmt.Println("Sending mail via", mail.Config().Host)
		err := mail.Send("Test")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("OK")
}
