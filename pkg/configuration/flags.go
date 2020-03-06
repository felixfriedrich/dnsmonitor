package configuration

import (
	"dnsmonitor/pkg/alerting"
	"errors"
	"flag"
)

// Flags contains values parsed from command line flags
type Flags struct {
	Domains    Domains
	DNS        string
	Silent     bool
	Interval   int
	Mail       bool
	SMS        bool
	Version    bool
	ConfigFile string
}

// ParseFlags parses all command line flags and returns an object containing the values
func ParseFlags() Flags {
	f := Flags{}
	flag.Var(&f.Domains, "domain", "domain to check. Can be used multiple times.")
	flag.StringVar(&f.DNS, "dns", "8.8.8.8", "DNS server")
	flag.BoolVar(&f.Silent, "silent", false, "silence output")
	flag.IntVar(&f.Interval, "interval", 300, "interval in seconds")
	flag.BoolVar(&f.Mail, "mail", false, "send mail if DNS record changes")
	flag.BoolVar(&f.SMS, "sms", false, "send SMS if DNS record changes")
	flag.BoolVar(&f.Version, "version", false, "print version")
	flag.StringVar(&f.ConfigFile, "configfile", "", "config file")
	flag.Parse()
	return f
}

func createConfigFromFlags(flags Flags) Config {
	configMap := make(Config)
	configMap["default"] = Check{
		Domains:  flags.Domains,
		DNS:      flags.DNS,
		Silent:   flags.Silent,
		Interval: flags.Interval,
		Mail:     flags.Mail,
		SMS:      flags.SMS,
	}
	return configMap
}

// Domains is a list of domains specified via command line or config file
type Domains []string

// String satisfies flag.Value (needed for flag parsing). See: https://golang.org/pkg/flag/#Value
func (d *Domains) String() string {
	r := "["
	for i, s := range *d {
		r = r + s
		if i != len(*d)-1 {
			r = r + ", "
		}
	}
	return r + "]"
}

// Set satisfies flag.Value (needed for flag parsing). See: https://golang.org/pkg/flag/#Value
func (d *Domains) Set(value string) error {
	*d = append(*d, value)
	return nil
}

// VendorFlag contains a validated vendor
type VendorFlag struct {
	Vendor alerting.Vendor
}

// TODO: What map ca be used to lookup keys and values? Use it for Set() and String()
func (vf *VendorFlag) String() string {
	if vf.Vendor == alerting.None {
		return "none"
	}
	if vf.Vendor == alerting.MessageBird {
		return "messagebird"
	}
	if vf.Vendor == alerting.SMS77 {
		return "sms77"
	}
	panic("")
}

// Set satisfies flag.Value
func (vf *VendorFlag) Set(f string) error {
	if f == "none" {
		vf.Vendor = alerting.None
		return nil
	}
	if f == "messagebird" {
		vf.Vendor = alerting.MessageBird
		return nil
	}
	if f == "sms77" || f == "SMS77" {
		vf.Vendor = alerting.SMS77
		return nil
	}
	return errors.New("vendor flag unknown")
}
