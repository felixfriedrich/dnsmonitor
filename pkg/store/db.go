package store

var (
	data map[string]Domain
)

func init() {
	data = map[string]Domain{}
}

// CleanUp would clean files from disk in the future
func CleanUp() error {
	return nil
}

// Get fetches an object for that domain or returns a new entry in case there isn't any yet
func Get(domain string) (Domain, error) {
	if d, ok := data[domain]; ok {
		return d, nil
	}
	return CreateDomain(domain), nil
}

// Save saves a domain object
func Save(domain Domain) error {
	data[domain.Name] = domain
	return nil
}
