package store

import (
	"dnsmonitor/pkg/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	r := m.Run()
	CleanUp()
	os.Exit(r)
}

func TestSave(t *testing.T) {
	d := model.CreateDomain("example.com")
	err := Save(d)
	assert.NoError(t, err)
}

func TestGetNewDomain(t *testing.T) {
	d, err := Get("example.com")
	assert.NoError(t, err)
	assert.Equal(t, "example.com", d.Name)
	assert.Len(t, d.Observations, 0)
}

func TestGetExisting(t *testing.T) {
	d := model.CreateDomain("example.com")
	d.Observations = append(d.Observations, *model.CreateRecord([]string{"93.184.216.34"}))
	Save(d)

	g, err := Get("example.com")
	assert.NoError(t, err)
	assert.Len(t, g.Observations, 1)
}
