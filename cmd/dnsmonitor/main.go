package main

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/dns"
	"dnsmonitor/pkg/monitor"
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

func main() {
	flags := config.ParseFlags()

	if flags.Version {
		fmt.Println("dnsmonitor v0.2")
		os.Exit(0)
	}

	monitors := []monitor.Monitor{}
	for _, d := range flags.Domains {
		configuration := config.CreateConfigFromFlags(flags)

		var alertingAPI alerting.API
		if configuration.SMS {
			alertingAPI = alerting.New(alerting.MessageBird, alerting.SMS)
		}

		var mail alerting.Mail
		if configuration.Mail {
			c := alerting.MailConfig{}
			prefix := "dnsmonitor_mail"
			err := envconfig.Process(prefix, &c)
			config.HandleEnvConfigError(err, c, prefix)
			mail = alerting.CreateMail(c)
		}

		m, err := monitor.CreateMonitor(d, configuration, mail, alertingAPI, dns.New())
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
