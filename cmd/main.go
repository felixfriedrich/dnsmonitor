package main

import (
	"dnsmonitor/pkg/store"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
)

func checkDomain(domain string, silent bool) {
	m := dns.Msg{}
	m.SetQuestion(domain+".", dns.TypeA)
	dnsClient := dns.Client{}
	r, t, err := dnsClient.Exchange(&m, "8.8.8.8:53")
	if !silent {
		fmt.Println("DNS query took", t)
	}
	if err != nil {
		log.Fatal(err)
	}

	answers := []string{}
	for _, a := range r.Answer {
		answers = append(answers, strings.Fields(a.String())[4])
	}

	d, err := store.Get(domain)
	if err != nil {
		log.Fatal(err)
	}
	d.Observations = append(d.Observations, store.CreateRecord(answers))
	store.Save(d)

	fmt.Println("Found", len(r.Answer), "answer(s).")
	for _, a := range answers {
		fmt.Println(a)
	}
}

func main() {
	var domain string
	flag.StringVar(&domain, "domain", "", "domain")
	var silent bool
	flag.BoolVar(&silent, "silent", false, "silence output")

	flag.Parse()

	if !silent {
		fmt.Println("Checking domain", domain)
	}

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			checkDomain(domain, silent)
		}
	}
}
