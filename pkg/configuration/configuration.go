package configuration

// Config holds configuration.
type Config struct {
	Domains  Domains
	DNS      string
	Silent   bool
	Interval int
	Mail     bool
	SMS      bool
}

// FromFlags takes command line flags and converts them into a generic Config object
func FromFlags(flags Flags) Config {
	return Config{
		Domains:  flags.Domains,
		DNS:      flags.DNS,
		Silent:   flags.Silent,
		Interval: flags.Interval,
		Mail:     flags.Mail,
		SMS:      flags.SMS,
	}
}
