package alerting

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/configuration/envconfig"
	"net/smtp"
	"strconv"
)

// Mail interface for mocking
//counterfeiter:generate . Mail
type Mail interface {
	Send(diff string) error
	Config() configuration.MailAlerting
}

type mail struct {
	config configuration.MailAlerting
}

func newMailFromConfig(config configuration.MailAlerting) Mail {
	return &mail{config: config}
}

// NewMail returns a mail implementation satisfying the interface alerting.Mail
func NewMail() Mail {
	c := configuration.MailAlerting{}
	prefix := "dnsmonitor_mail"
	envconfig.Read(prefix, &c)
	return newMailFromConfig(c)
}

func (m *mail) Config() configuration.MailAlerting {
	return m.config
}

// Send sends e-mails
func (m *mail) Send(diff string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		m.config.Username,
		m.config.Password,
		m.config.Host,
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		m.config.Host+":"+strconv.Itoa(m.config.Port),
		auth,
		m.config.From,
		[]string{m.config.To},
		[]byte(diff),
	)
	if err != nil {
		return err
	}
	return nil
}
