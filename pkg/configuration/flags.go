package configuration

import (
	"errors"
	"flag"
)

const (
	// None is used as default by the 'flags' package
	None Vendor = "" // This needs to be an empty string in order to match the default value from yml parsing
	// MessageBird https://www.messagebird.com/en/
	MessageBird Vendor = "messagebird"
	// SMS77 https://app.sms77.io
	SMS77 Vendor = "sms77"
)

// Vendor identifies services for alerting
type Vendor string

// Flags contains values parsed from command line flags
type Flags struct {
	Domains    Domains
	DNS        string
	Silent     bool
	Interval   int
	Mail       bool
	SMS        bool
	VendorFlag VendorFlag
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
	flag.Var(&f.VendorFlag, "vendor", "Alerting vendor, e.g. 'messagebird'.")
	flag.BoolVar(&f.Version, "version", false, "print version")
	flag.StringVar(&f.ConfigFile, "configfile", "", "config file")
	flag.Parse()
	return f
}

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
	Vendor Vendor
}

func (vf *VendorFlag) String() string {
	return string(vf.Vendor)
}

// Set satisfies flag.Value
func (vf *VendorFlag) Set(f string) error {
	vendors := map[Vendor]bool{MessageBird: true, SMS77: true} // existing vendors can be disabled via this map
	enabled, exists := vendors[Vendor(f)]
	if exists && enabled {
		vf.Vendor = Vendor(f)
		return nil
	}
	return errors.New("vendor doesn't exist")
}
