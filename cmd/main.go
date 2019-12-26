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
		fmt.Println("dnsmonitor v0.2")
		os.Exit(0)
	}

	monitors := []dns.Monitor{}
	for _, d := range flags.Domains {
		m, err := dns.CreateMonitor(d, config.CreateConfigFromFlags(flags))
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
