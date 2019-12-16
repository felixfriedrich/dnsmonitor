package store

import "dnsmonitor/pkg/model"

var (
	data map[string]*model.Domain
)

func init() {
	data = map[string]*model.Domain{}
}

// CleanUp would clean files from disk in the future
func CleanUp() error {
	return nil
}

// Get fetches an object for that domain or returns a new entry in case there isn't any yet
func Get(domain string) (*model.Domain, error) {
	if d, ok := data[domain]; ok {
		return d, nil
	}
	return model.CreateDomain(domain), nil
}

// Save saves a domain object
func Save(domain *model.Domain) error {
	data[domain.Name] = domain
	return nil
}
