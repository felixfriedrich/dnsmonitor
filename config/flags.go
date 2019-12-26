package config

import (
	"flag"
)

type domains []string

func (d *domains) String() string {
	r := "["
	for i, s := range *d {
		r = r + s
		if i != len(*d)-1 {
			r = r + ", "
		}
	}
	return r + "]"
}

func (d *domains) Set(value string) error {
	*d = append(*d, value)
	return nil
}

// Flags contains values parsed from command line flags
type Flags struct {
	Domains  domains
	DNS      string
	Silent   bool
	Interval int
	Mail     bool
	SMS      bool
	Version  bool
}

// ParseFlags parses all command line flags and returns an object containing the values
func ParseFlags() Flags {
	f := Flags{}
	flag.Var(&f.Domains, "domain", "domain to check. Can be used multiple times.")
	flag.StringVar(&f.DNS, "dns", "8.8.8.8", "DNS server")
	flag.BoolVar(&f.Silent, "silent", false, "silence output")
	flag.IntVar(&f.Interval, "interval", 1, "interval in seconds")
	flag.BoolVar(&f.Mail, "mail", false, "send mail if DNS record changes")
	flag.BoolVar(&f.SMS, "sms", false, "send SMS if DNS record changes")
	flag.BoolVar(&f.Version, "version", false, "print version")
	flag.Parse()
	return f
}
