package configuration

import (
	"dnsmonitor/pkg/alerting"
	"errors"
	"flag"
)

// Flags contains values parsed from command line flags
type Flags struct {
	Domains  DomainFlag
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

// DomainFlag is a list of domain specified via command line
type DomainFlag []string

func (d *DomainFlag) String() string {
	r := "["
	for i, s := range *d {
		r = r + s
		if i != len(*d)-1 {
			r = r + ", "
		}
	}
	return r + "]"
}

// Set satisfies flag.Value
func (d *DomainFlag) Set(value string) error {
	*d = append(*d, value)
	return nil
}

// VendorFlag contains a validated vendor
type VendorFlag struct {
	Vendor alerting.Vendor
}

func (vf VendorFlag) String() string {
	if vf.Vendor == alerting.None {
		return "none"
	}
	if vf.Vendor == alerting.MessageBird {
		return "messagebird"
	}
	panic("")
}

// Set satisfies flag.Value
func (vf VendorFlag) Set(f string) error {
	if f == "none" {
		vf.Vendor = alerting.None
		return nil
	}
	if f == "messagebird" {
		vf.Vendor = alerting.MessageBird
		return nil
	}
	return errors.New("vendor flag unknown")
}
