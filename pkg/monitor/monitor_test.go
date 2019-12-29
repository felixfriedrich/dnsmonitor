package monitor

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/alerting/alertingfakes"
	"dnsmonitor/pkg/dns/dnsfakes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	c := config.Config{
		Domains:  []string{"google.com", "www.google.com"},
		DNS:      "8.8.8.8",
		Silent:   false,
		Interval: 300,
		Mail:     false,
	}
	dns := &dnsfakes.FakeInterface{}
	dns.QueryReturnsOnCall(0, []string{"1.2.3.4"}, nil)
	m, err := CreateMonitor("www.google.com", c, nil, nil, dns)
	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.NotEmpty(t, m.Domain())
	assert.NotNil(t, m.Config().Domains)
	m.Check()
	assert.Equal(t, 1, dns.QueryCallCount())
	domain, _ := dns.QueryArgsForCall(0)

	assert.Equal(t, domain, "www.google.com")
}

func TestCreateMonitorWithSMSAlerting(t *testing.T) {
	c := config.Config{
		Domains:  []string{"google.com", "www.google.com"},
		DNS:      "8.8.8.8",
		Silent:   false,
		Interval: 300,
		Mail:     false,
		SMS:      true,
	}

	alertingAPI := &alertingfakes.FakeAPI{}
	dns := &dnsfakes.FakeInterface{}
	dns.QueryReturnsOnCall(0, []string{"1.2.3.4"}, nil)
	m, err := CreateMonitor("www.google.com", c, nil, alertingAPI, dns)
	m.Check()
	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.Equal(t, 1, alertingAPI.SendSMSCallCount())
}

func TestCreateMonitorWithMailAlerting(t *testing.T) {
	c := config.Config{
		Domains:  []string{"google.com", "www.google.com"},
		DNS:      "8.8.8.8",
		Silent:   false,
		Interval: 300,
		Mail:     true,
		SMS:      false,
	}
	dns := &dnsfakes.FakeInterface{}
	dns.QueryReturnsOnCall(0, []string{"1.2.3.4"}, nil)
	dns.QueryReturnsOnCall(1, []string{"4.3.2.1"}, nil)
	mail := &alertingfakes.FakeMail{}
	m, err := CreateMonitor("www.google.com", c, mail, nil, dns)
	assert.NoError(t, err)
	assert.NotNil(t, m)
	m.Check()
	assert.Equal(t, 0, mail.SendCallCount())
	m.Check()
	assert.Equal(t, 1, mail.SendCallCount())
}
