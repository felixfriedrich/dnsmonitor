package main

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/dns"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	flags := config.ParseFlags()

	if flags.Version {
		fmt.Println("dnsmonitor v0.1")
		os.Exit(0)
	}

	if !flags.Silent {
		fmt.Println("Checking domain", flags.Domain)
	}

	ticker := time.NewTicker(time.Duration(flags.Interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			m, err := dns.CreateMonitor(flags)
			if err != nil {
				log.Error(err)
			}
			m.Check()
		}
	}
}
