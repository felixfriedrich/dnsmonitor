package main

import (
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/dns"
	"dnsmonitor/pkg/monitor"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	flags := configuration.ParseFlags()

	if flags.Version {
		fmt.Println("dnsmonitor v0.3")
		os.Exit(0)
	}

	monitors := []monitor.Monitor{}
	config := configuration.Create(flags)

	var alertingAPI alerting.API
	var err error
	if config.SMS {
		alertingAPI, err = alerting.New(alerting.MessageBird, alerting.SMS)
		if err != nil {
			log.Fatal(err)
		}
	}

	var mail alerting.Mail
	if config.Mail {
		mail = alerting.NewMail()
	}

	m, err := monitor.CreateMonitor(config, mail, alertingAPI, dns.New())
	if err != nil {
		log.Error(err)
	}
	monitors = append(monitors, m)

	ticker := time.NewTicker(time.Duration(config.Interval) * time.Second)
	for {
		select {
		case <-ticker.C:

			m.Check()

			for _, d := range m.Domains() {
				if !config.Silent {
					fmt.Println("Checking domain", d.Name)
				}

				fmt.Println("Found", len(d.LastObservation().GetAnswers()), "answer(s).")
				for _, aa := range d.LastObservation().GetAnswers() {
					fmt.Println(aa)
				}

				diff, _ := d.GetDiff()
				fmt.Println(diff)
			}
		}
	}
}
