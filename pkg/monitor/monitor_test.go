package monitor

import (
	"dnsmonitor/pkg/alerting/alertingfakes"
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/dns/dnsfakes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	c := configuration.Config{
		Domains: []string{"google.com", "www.google.com"},
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

func TestCreateMonitorWithAlerting(t *testing.T) {
	c := configuration.Config{
		Mail: true,
		SMS:  true,
	}
	dns := &dnsfakes.FakeInterface{}
	dns.QueryReturnsOnCall(0, []string{"1.2.3.4"}, nil)
	dns.QueryReturnsOnCall(1, []string{"4.3.2.1"}, nil)
	mail := &alertingfakes.FakeMail{}
	alertingAPI := &alertingfakes.FakeAPI{}
	m, err := CreateMonitor("www.google.com", c, mail, alertingAPI, dns)
	assert.NoError(t, err)
	assert.NotNil(t, m)
	m.Check()
	assert.Equal(t, 0, mail.SendCallCount())
	assert.Equal(t, 0, alertingAPI.SendSMSCallCount())
	m.Check()
	assert.Equal(t, 1, mail.SendCallCount())
	assert.Equal(t, 1, alertingAPI.SendSMSCallCount())
}

