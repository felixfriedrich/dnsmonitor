package messagebird

import (
	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

// MessageBird holds a client to connect to the messagebird api
type messageBird struct {
	Client *messagebird.Client
	Config Config
}

// Config for envconfig
type Config struct {
	AccessKey   string   `required:"true"`
	Sender      string   `required:"true"`
	Receipients []string `required:"true"`
}

// New creates a messageBird struct
func New(config Config) *messageBird {
	return &messageBird{Client: messagebird.New(config.AccessKey), Config: config}
}

func (mb *messageBird) SendSMS(text string) error {
	_, err := sms.Create(mb.Client, mb.Config.Sender, mb.Config.Receipients, text, nil)
	if err != nil {
		return err
	}
	return nil
}
