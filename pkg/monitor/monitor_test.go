package monitor

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/alerting/alertingfakes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMonitor(t *testing.T) {
	c := config.Config{
		Domains:  []string{"google.com", "www.google.com"},
		DNS:      "8.8.8.8",
		Silent:   false,
		Interval: 300,
		Mail:     false,
	}
	m, err := CreateMonitor("www.google.com", c, nil)
	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.NotEmpty(t, m.Domain())
	assert.NotNil(t, m.Config().Domains)
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
	m, err := CreateMonitor("www.google.com", c, alertingAPI)
	m.Check()
	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.Equal(t, 1, alertingAPI.SendSMSCallCount())
}
