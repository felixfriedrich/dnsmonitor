package sms77

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"dnsmonitor/pkg/configuration"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// SMS77 holds an HTTP client and some config information
type SMS77 struct {
	Client HTTPClient
	Config configuration.SMS77Config
}

// Override sets a new http.Client into the SMS77 struct (used for testing)
func (s *SMS77) Override(httpClient HTTPClient) {
	s.Client = httpClient
}

// HTTPClient encapsulates the http library for testing/mocking
//
//counterfeiter:generate . HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New created a SMS77 object from a config
func New(config configuration.SMS77Config) *SMS77 {
	return &SMS77{Client: &http.Client{}, Config: config}
}

// SendSMS satisfies the alerting.API interface
// https://www.sms77.io/en/docs/gateway/http-api/sms-disptach/
func (s *SMS77) SendSMS(text string) error {
	url := fmt.Sprintf("https://gateway.sms77.io/api/sms?%s", createURLValues(s, text).Encode())
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// https://www.sms77.io/en/docs/gateway/http-api/general/
	request.Header.Add("Authorization", fmt.Sprintf("basic %s", s.Config.APIKey))

	response, err := s.Client.Do(request)
	if err != nil {
		return err
	}

	// Checking HTTP status code
	log.Info("http status code is ", response.StatusCode)
	if response.StatusCode != 200 {
		return errors.New("sms77 returned non-200 status code")
	}

	log.Debug(requestHeaders(request))
	log.Debug(responseHeaders(response))

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	log.Debug("Response body", string(body))

	result := parseResponseBody(body)
	log.Debug(result)

	// Checking SMS77 status code
	// https://www.sms77.io/en/docs/gateway/http-api/sms-disptach
	if result.Success == "100" || result.Success == "102" {
		return nil
	}

	log.Error("sms77 returned status code ", result.Success)
	return errors.New("sms77 returned unsuccessfully")
}

func parseResponseBody(body []byte) Response {
	var result Response
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func requestHeaders(request *http.Request) string {
	requestHeaders, err := httputil.DumpRequestOut(request, false)
	if err != nil {
		log.Fatalln(err)
	}
	return string(requestHeaders)
}

func responseHeaders(response *http.Response) string {
	responseHeaders, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	return string(responseHeaders)
}

func createURLValues(s *SMS77, text string) url.Values {
	payload := url.Values{}
	payload.Set("to", s.Config.Recipient)
	payload.Set("type", "direct")
	payload.Set("text", text)
	payload.Set("from", s.Config.Sender)
	payload.Set("json", "1")

	if s.Config.Debug {
		payload.Set("debug", "1")
	}
	return payload
}

// Response is used to parse the JSON response from SMS77
type Response struct {
	Balance  float32   `json:"balance"`
	Debug    string    `json:"debug"`
	Price    float32   `json:"total_price"`
	Success  string    `json:"success"`
	Messages []Message `json:"messages"`
}

// Message is used to parse the JSON response from SMS77
type Message struct {
	Success bool `json:"success"`
}

func (r Response) String() string {
	s := fmt.Sprintf("Success: %s (See: https://www.sms77.io/en/docs/gateway/http-api/sms-disptach/)\n", r.Success)
	s += fmt.Sprintf("Balance: %f EUR\n", r.Balance)
	s += fmt.Sprintf("Messages (%d) %s", len(r.Messages), r.Messages)
	return s
}

func (m Message) String() string {
	return fmt.Sprintf("Success: %t", m.Success)
}
