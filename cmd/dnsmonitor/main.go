package main

import (
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/dns"
	"dnsmonitor/pkg/monitor"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	version                         = "0.4.2"
	okExitCode                      = 0
	fileDoesntExistExitCode         = 1
	wrongCombinationOfFlagsExitCode = 2
)

func main() {
	flags := configuration.ParseFlags()

	if flags.Version {
		fmt.Printf("dnsmonitor v%s\n", version)
		os.Exit(okExitCode)
	}

	if flags.ConfigFile != "" {
		_, err := os.Stat(flags.ConfigFile)
		if os.IsNotExist(err) {
			fmt.Println("config file", flags.ConfigFile, "doesn't exist")
			os.Exit(fileDoesntExistExitCode)
		}
	}

	if flags.ConfigFile != "" && len(flags.Domains) > 0 {
		fmt.Println("Provide either -configfile OR -domain")
		os.Exit(wrongCombinationOfFlagsExitCode)
	}

	for _, config := range configuration.CreateConfig(flags).Monitors {
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

		go m.Run(config.Interval, config.Silent)
	}

	select {} // Make this program not terminate in order to keep the go routines running
}
