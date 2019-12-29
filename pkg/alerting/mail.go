package alerting

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"net/smtp"
	"strconv"
)

// Mail interface for mocking
//counterfeiter:generate . Mail
type Mail interface {
	Send(diff string) error
}

type mail struct {
	Config MailConfig
}

// CreateMail creates new Mail
func CreateMail(config MailConfig) Mail {
	return &mail{Config: config}
}

// Send sends e-mails
func (m *mail) Send(diff string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		m.Config.Username,
		m.Config.Password,
		m.Config.Host,
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		m.Config.Host+":"+strconv.Itoa(m.Config.Port),
		auth,
		m.Config.From,
		[]string{m.Config.To},
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
