package main

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/store"
	"fmt"
	"net/smtp"
	"strconv"
	"time"

	"github.com/pmezard/go-difflib/difflib"
	log "github.com/sirupsen/logrus"
)

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

func sendMail(diff string) error {
	c := config.CreateMailConfigFromEnvOrDie()
	if diff != "" {
		// Set up authentication information.
		auth := smtp.PlainAuth(
			"",
			c.Username,
			c.Password,
			c.Host,
		)
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		err := smtp.SendMail(
			c.Host+":"+strconv.Itoa(c.Port),
			auth,
			c.From,
			[]string{c.To},
			[]byte(diff),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flags := config.ParseFlags()

	if !flags.Silent {
		fmt.Println("Checking domain", flags.Domain)
	}

	ticker := time.NewTicker(time.Duration(flags.Interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			d, err := store.Get(flags.Domain)
			if err != nil {
				log.Fatal(err)
			}

			r := d.Observe()
			if !flags.Silent {
				fmt.Println("Found", len(r.GetAnswers()), "answer(s).")
				for _, aa := range r.GetAnswers() {
					fmt.Println(aa)
				}
			}

			diff, err := getDiff(d, r.GetAnswers())
			if err != nil {
				log.Error(err)
			}
			if !flags.Silent {
				fmt.Println(diff)
			}
			if flags.Mail {
				err = sendMail(diff)
				if err != nil {
					log.Error(err)
				}
			}

			err = store.Save(d)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
