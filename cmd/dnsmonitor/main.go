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
		fmt.Println("dnsmonitor v0.2")
		os.Exit(0)
	}

	monitors := []monitor.Monitor{}
	for _, d := range flags.Domains {
		config := configuration.FromFlags(flags)

		var alertingAPI alerting.API
		if config.SMS {
			alertingAPI = alerting.New(alerting.MessageBird, alerting.SMS)
		}

		var mail alerting.Mail
		if config.Mail {
			mail = alerting.NewMail()
		}

		m, err := monitor.CreateMonitor(d, config, mail, alertingAPI, dns.New())
		if err != nil {
			log.Error(err)
		}
		monitors = append(monitors, m)
	}

	ticker := time.NewTicker(time.Duration(flags.Interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			for _, m := range monitors {
				if !flags.Silent {
					fmt.Println("Checking domain", m.Domain())
				}
				r := m.Check()

				if !flags.Silent {
					fmt.Println("Found", len(r.GetAnswers()), "answer(s).")
					for _, aa := range r.GetAnswers() {
						fmt.Println(aa)
					}

					diff, _ := m.Domain().GetDiff()
					fmt.Println(diff)
				}
			}
		}
	}
}
