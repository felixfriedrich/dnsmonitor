package alerting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendSMSViaMessageBird(t *testing.T) {
	s, err := MessageBird()
	// err := config.CreateEnvConfigFromEnvOrDie(&c, "dnsmonitor_messagebird")
	// This test can only be run with a config.
	// In case there isn't any config, this test has to be skipped. e.g. in CI/CD.
	// Note: for this reason this isn't a real unit test, but more like an integration test.
	if err != nil {
		t.SkipNow()
	}
	assert.NotNil(t, s)
	msg, err := s.Send("Test")
	assert.NoError(t, err)
	assert.NotNil(t, msg)
}
