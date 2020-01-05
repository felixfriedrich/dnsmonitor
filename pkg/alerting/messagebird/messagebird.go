package messagebird

import (
	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

// MessageBird holds a client to connect to the messagebird api
// This struct could be private as its never used directly, only via the interface alerting.API
// Yet, the linter isn't happy with it being private :-/
type MessageBird struct {
	Client *messagebird.Client
	Config Config
}

// Config for envconfig
type Config struct {
	AccessKey  string   `required:"true"`
	Sender     string   `required:"true"`
	Recipients []string `required:"true"`
}

// New creates a messageBird struct
func New(config Config) *MessageBird {
	return &MessageBird{Client: messagebird.New(config.AccessKey), Config: config}
}

// SendSMS satisfies the interface alerting.API
func (mb *MessageBird) SendSMS(text string) error {
	_, err := sms.Create(mb.Client, mb.Config.Sender, mb.Config.Recipients, text, nil)
	if err != nil {
		return err
	}
	return nil
}
