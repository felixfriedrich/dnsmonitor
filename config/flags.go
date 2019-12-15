package config

import "flag"

// Flags contains values parsed from command line flags
type Flags struct {
	Domain   string
	DNS      string
	Silent   bool
	Interval int
	Mail     bool
	Version  bool
}

// ParseFlags parses all command line flags and returns an object containing the values
func ParseFlags() Flags {
	f := Flags{}
	flag.StringVar(&f.Domain, "domain", "", "domain")
	flag.StringVar(&f.DNS, "DNS", "8.8.8.8", "DNS server")
	flag.BoolVar(&f.Silent, "silent", false, "silence output")
	flag.IntVar(&f.Interval, "interval", 1, "interval in seconds")
	flag.BoolVar(&f.Mail, "mail", false, "send mail if DNS record changes")
	flag.BoolVar(&f.Version, "version", false, "print version")
	flag.Parse()
	return f
}
