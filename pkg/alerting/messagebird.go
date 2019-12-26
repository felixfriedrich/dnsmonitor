// Run like this: DNSMONITOR_MESSAGEBIRD_ACCESSKEY="" go test ./pkg/alerting/

package alerting

import (
	"dnsmonitor/config"

	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

var (
	prefix = "dnsmonitor_mail"
)

type messageBird struct {
	Client *messagebird.Client
	Config MessageBirdConfig
}

// MessageBird provides SMS sending via MessageBird (https://developers.messagebird.com)
func MessageBird() (SMS, error) {
	c := MessageBirdConfig{}
	err := config.CreateEnvConfigFromEnv("dnsmonitor_messagebird", &c)
	if err != nil {
		return nil, err
	}

	client := messagebird.New(c.AccessKey)
	return &messageBird{Client: client, Config: c}, nil
}

func (m messageBird) Send(message string) (interface{}, error) {
	msg, err := sms.Create(m.Client, m.Config.Sender, m.Config.Recipients, message, nil)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// MessageBirdConfig defines environment variables to configure the MessageBird service
type MessageBirdConfig struct {
	AccessKey  string   `required:"true"`
	Sender     string   `required:"true"`
	Recipients []string `required:"true"`
}
