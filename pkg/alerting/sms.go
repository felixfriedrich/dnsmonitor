package alerting

// SMS is a generic interface for different SMS vendors
type SMS interface {
	Send(msg string) (interface{}, error)
}
