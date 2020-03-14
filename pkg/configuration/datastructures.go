package configuration

// NewConfig creates a new Config object and initialised the list of monitors
func NewConfig() Config {
	monitors := make(monitors)
	return Config{Monitors: monitors}
}

// Config represents the yml structure expected in a configuration file
type Config struct {
	Monitors monitors `yaml:"monitors"`
}

type monitors map[string]*Monitor

// Monitor is part of the yml structure expected in a configuration file
type Monitor struct {
	Domains  Domains  `yaml:"domains"`
	DNS      string   `yaml:"dns"`
	Interval int      `yaml:"interval"`
	Mail     bool     `yaml:"mail"`
	SMS      bool     `yaml:"sms"`
	Silent   bool     `yaml:"silent"`
	Alerting Alerting `yaml:"alerting"`
}

// Domains is a list of domains specified via command line or config file
type Domains []string

// Alerting hold information for alerting
type Alerting struct {
	Mail MailAlerting `yaml:"mail"`
	SMS  SMSAlerting  `yaml:"sms"`
}

// MailAlerting holds information for alerting via mail
type MailAlerting struct {
	Host     string `yaml:"host" default:"127.0.0.1"`
	Port     int    `yaml:"port" default:"25"`
	Username string `yaml:"username" required:"true"`
	Password string `yaml:"password" required:"true"`
	From     string `yaml:"from" required:"true"`
	To       string `yaml:"to" required:"true"`
}

// SMSAlerting holds information for alerting via Messagebird
type SMSAlerting struct {
	Vendor      Vendor            `yaml:"vendor"`
	MessageBird MessageBirdConfig `yaml:"messagebird"`
	SMS77       SMS77Config       `yaml:"sms77"`
}

// MessageBirdConfig for envconfig and yml files
type MessageBirdConfig struct {
	AccessKey  string   `yaml:"accesskey" required:"true"`
	Sender     string   `yaml:"sender" required:"true"`
	Recipients []string `yaml:"recipients" required:"true"`
}

// SMS77Config for envconfig and yml files
type SMS77Config struct {
	APIKey    string `required:"true" yaml:"apikey"`
	Sender    string `required:"true" yaml:"sender"`
	Recipient string `required:"true" yaml:"recipient"`
	Debug     bool   `required:"false" default:"false" yaml:"debug"`
}
