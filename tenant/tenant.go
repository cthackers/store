package tenant

type Tenant struct {
	name string
}

func (t *Tenant) Authenticate() *Tenant {
	return nil
}

func (t *Tenant) Close() error {
	return nil
}

func (t *Tenant) Collections() []string {
	return nil
}

func (t *Tenant) DeleteCollection(name string) error {
	return nil
}
