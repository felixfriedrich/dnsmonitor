package main

import (
	"dnsmonitor/pkg/configuration"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanityCheckFlags_Version(t *testing.T) {
	flags := configuration.Flags{
		Version: true,
	}

	ok, exitCode := sanityCheckFlags(flags)
	assert.False(t, ok)
	assert.Equal(t, okExitCode, exitCode)
}

func TestSanityCheckFlags_ConfigFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "dnsmonitor-test")
	if err != nil {
		log.Fatal(err)
	}
	os.Remove(tempFile.Name()) // make sure the file doesn't exist

	flags := configuration.Flags{
		ConfigFile: tempFile.Name(),
	}

	ok, exitCode := sanityCheckFlags(flags)
	assert.False(t, ok)
	assert.Equal(t, fileDoesntExistExitCode, exitCode)
}

func TestSanityCheckFlags_ConfigFileAndDomains(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "dnsmonitor-test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempFile.Name()) // eventually clean up

	flags := configuration.Flags{
		ConfigFile: tempFile.Name(),
		Domains:    configuration.Domains{"google.com"},
	}

	ok, exitCode := sanityCheckFlags(flags)
	assert.False(t, ok)
	assert.Equal(t, wrongCombinationOfFlagsExitCode, exitCode)
}
