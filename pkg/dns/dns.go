package dns

import (
	"time"

	lib "github.com/miekg/dns"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// Interface abtracts the upstream library away, mainly for mocking
//counterfeiter:generate . Interface
type Interface interface {
	Exchange(m *lib.Msg, address string) (r *lib.Msg, rtt time.Duration, err error)
}

type d struct {
	client lib.Client
}

// New returns an interface to the DNS library
func New() Interface {
	return &d{client: lib.Client{}}
}

func (d *d) Exchange(m *lib.Msg, address string) (r *lib.Msg, rtt time.Duration, err error) {
	return d.client.Exchange(m, address)
}
