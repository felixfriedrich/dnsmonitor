package main

import (
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/configuration/envconfig"
	"flag"
	"fmt"
)

type flags struct {
	VendorFlag configuration.VendorFlag
	SMS        bool
	Mail       bool
}

func main() {
	f := flags{}
	flag.Var(&f.VendorFlag, "vendor", "Alerting vendor, e.g. 'messagebird'.")
	flag.BoolVar(&f.SMS, "sms", false, "Test SMS alerting")
	flag.BoolVar(&f.Mail, "mail", false, "Test mail alerting")
	flag.Parse()

	if f.SMS {
		fmt.Println("Using SMS vendor:", f.VendorFlag.String())
		alertingAPI, err := alerting.New(f.VendorFlag.Vendor, alerting.SMS)
		if err != nil {
			panic(err)
		}
		err = alertingAPI.SendSMS("Test")
		if err != nil {
			panic(err)
		}
	}

	if f.Mail {
		ma := configuration.MailAlerting{}
		envconfig.Read(configuration.EnvMailPrefix, &ma)
		mail := alerting.NewMailAlerting(ma)
		fmt.Println("Sending mail via", mail.Config().Host)
		err := mail.Send("Test")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("OK")
}
