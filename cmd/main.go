package main

import (
	"dnsmonitor/pkg/store"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/miekg/dns"
	"github.com/pmezard/go-difflib/difflib"
	log "github.com/sirupsen/logrus"
)

func checkDomain(domain string, silent bool) []string {
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

	return answers
}

func getDiff(domain store.Domain, answers []string) (string, error) {
	lo := domain.GetLastObservation()
	diff := difflib.UnifiedDiff{
		A:        lo.GetAnswers(),
		B:        answers,
		FromFile: "Last answer " + lo.Time().String(),
		ToFile:   "Current answer " + time.Now().String(),
		Context:  1,
	}
	return difflib.GetUnifiedDiffString(diff)
}

func main() {
	var domain string
	flag.StringVar(&domain, "domain", "", "domain")
	var silent bool
	flag.BoolVar(&silent, "silent", false, "silence output")
	var interval int
	flag.IntVar(&interval, "interval", 1, "interval in seconds")

	flag.Parse()

	if !silent {
		fmt.Println("Checking domain", domain)
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			a := checkDomain(domain, silent)
			fmt.Println("Found", len(a), "answer(s).")
			for _, aa := range a {
				fmt.Println(aa)
			}

			d, err := store.Get(domain)
			if err != nil {
				log.Fatal(err)
			}

			diff, err := getDiff(d, a)
			if err != nil {
				log.Error(err)
			}
			fmt.Println(diff)

			d.Observations = append(d.Observations, store.CreateRecord(a))
			store.Save(d)
		}
	}
}
