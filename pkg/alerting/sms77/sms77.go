package sms77

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// SMS77 holds an HTTP client and some config information
type SMS77 struct {
	Client *http.Client
	Config Config
}

// New created a SMS77 object from a config
func New(config Config) *SMS77 {
	return &SMS77{Client: &http.Client{}, Config: config}
}

// Config for envconfig
type Config struct {
	APIKey    string `required:"true"`
	Sender    string `required:"true"`
	Recipient string `required:"true"`
	Debug     bool   `required:"false" default:"false"`
}

// SendSMS satifies the alerting.API interface
// https://www.sms77.io/en/docs/gateway/http-api/sms-disptach/
func (s *SMS77) SendSMS(text string) error {
	payload := url.Values{}
	payload.Set("to", s.Config.Recipient)
	payload.Set("type", "direct")
	payload.Set("text", text)
	payload.Set("from", s.Config.Sender)
	payload.Set("json", "1")

	if s.Config.Debug {
		payload.Set("debug", "1")
	}

	url := fmt.Sprintf("https://gateway.sms77.io/api/sms?%s", payload.Encode())
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

	requestHeaders, err := httputil.DumpRequestOut(request, false)
	if err != nil {
		log.Fatalln(err)
	}
	log.Debug(string(requestHeaders))

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	log.Debug("Response body", string(body))

	var result Response
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		panic(err)
	}
	log.Debug(result)

	// Checking SMS77 status code
	// https://www.sms77.io/en/docs/gateway/http-api/sms-disptach
	if result.Success == "100" || result.Success == "102" {
		return nil
	}

	log.Error("sms77 returned status code ", result.Success)
	return errors.New("sms77 returned unsuccessfully")
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
	// [map[encoding:gsm error:<nil> error_text:<nil> id:<nil> parts:1 price:0 recipient:491627062392 sender:491627062392 text:Test]]
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
