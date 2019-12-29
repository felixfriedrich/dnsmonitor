package alerting

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// API abtracts the upstream library away, mainly for mocking
//counterfeiter:generate . API
type API interface {
	SendSMS(text string) error
}
