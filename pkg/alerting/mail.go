package alerting

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"dnsmonitor/pkg/configuration/envconfig"
	"net/smtp"
	"strconv"
)

// Mail interface for mocking
//counterfeiter:generate . Mail
type Mail interface {
	Send(diff string) error
	Config() MailConfig
}

type mail struct {
	config MailConfig
}

func newMailFromConfig(config MailConfig) Mail {
	return &mail{config: config}
}

// NewMail returns a mail implementation satisfying the interface alerting.Mail
func NewMail() Mail {
	c := MailConfig{}
	prefix := "dnsmonitor_mail"
	envconfig.Read(prefix, &c)
	return newMailFromConfig(c)
}

func (m *mail) Config() MailConfig {
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

// MailConfig defines environment variables to configure the mail server to use
type MailConfig struct {
	Host     string `default:"127.0.0.1"`
	Port     int    `default:"25"`
	Username string `required:"true"`
	Password string `required:"true"`
	From     string `required:"true"`
	To       string `required:"true"`
}
