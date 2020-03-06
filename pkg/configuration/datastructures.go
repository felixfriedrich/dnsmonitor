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

type monitors map[string]Monitor

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

// Alerting hold information for alerting
type Alerting struct {
	Mail MailAlerting `yaml:"mail"`
}

// MailAlerting holds information for alerting via mail
type MailAlerting struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	To       string `yaml:"to"`
}
