package store

import (
	"github.com/cthackers/store/tenant"
)

type Options struct {
	Path        string
	CacheSize   int
	Compression bool
}

type Store struct {
	options Options
}

func New(options Options) *Store {
	store := &Store{}
	return store
}

func (s *Store) NewTenant() *tenant.Tenant {
	return nil
}

func (s *Store) GetTenant(id string) *tenant.Tenant {
	return nil
}
