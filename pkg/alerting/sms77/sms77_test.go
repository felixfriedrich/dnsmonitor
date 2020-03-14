package sms77_test

import (
	"dnsmonitor/pkg/alerting/sms77"
	"dnsmonitor/pkg/alerting/sms77/sms77fakes"
	"dnsmonitor/pkg/configuration"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMS77_SendSMS(t *testing.T) {
	api := sms77.New(configuration.SMS77Config{
		APIKey:    "1234",
		Sender:    "01234567890",
		Recipient: "0987654321",
		Debug:     true,
	})
	fakeHttpClient := sms77fakes.FakeHTTPClient{}
	fakeHttpClient.DoReturnsOnCall(0, &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader("{}")),
	}, nil)
	api.Override(&fakeHttpClient)
	api.SendSMS("Test")
	request := fakeHttpClient.DoArgsForCall(0)

	assert.Equal(t, "basic 1234", request.Header.Get("Authorization"))
	assert.Equal(t, "GET", request.Method)
	assert.Equal(t, "gateway.sms77.io", request.Host)

	params := url.Values{}
	params.Set("from", "01234567890")
	assert.Contains(t, request.URL.RawQuery, params.Encode())

	params = url.Values{}
	params.Set("to", "0987654321")
	assert.Contains(t, request.URL.RawQuery, params.Encode())

	params = url.Values{}
	params.Set("debug", "1")
	assert.Contains(t, request.URL.RawQuery, params.Encode())

	params = url.Values{}
	params.Set("text", "Test")
	assert.Contains(t, request.URL.RawQuery, params.Encode())

	assert.Equal(t, request.URL.Path, "/api/sms")
}

func TestSMS77_SendSMS_Non200Response(t *testing.T) {
	api := sms77.New(configuration.SMS77Config{})
	fakeHttpClient := sms77fakes.FakeHTTPClient{}
	fakeHttpClient.DoReturnsOnCall(0, &http.Response{
		StatusCode: 404,
		Body:       ioutil.NopCloser(strings.NewReader("{}")),
	}, nil)
	api.Override(&fakeHttpClient)
	err := api.SendSMS("Test")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sms77 returned non-200 status code")
}

func TestSMS77_SendSMS_SendFailed(t *testing.T) {
	api := sms77.New(configuration.SMS77Config{})
	fakeHttpClient := sms77fakes.FakeHTTPClient{}
	fakeHttpClient.DoReturnsOnCall(0, &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader("{\"Success\": \"101\"}")),
	}, nil)
	api.Override(&fakeHttpClient)
	err := api.SendSMS("Test")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sms77 returned unsuccessfully")
}
