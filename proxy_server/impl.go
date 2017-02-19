package proxy_server

import (
	"github.com/bcho/dmm"
	"github.com/bcho/dmm/search"
)

type serviceOpt func(s *service)

func WithSearchClient(client *search.Client) serviceOpt {
	return func(s *service) {
		s.client = client
	}
}

type service struct {
	client *search.Client
}

func NewService(opts ...serviceOpt) *service {
	s := &service{}

	for _, o := range opts {
		o(s)
	}

	if s.client == nil {
		panic("search client is required")
	}

	return s
}

func (s service) Query(keyword string) ([]dmm.Product, error) {
	return s.client.Query(&search.QueryOpts{keyword})
}

func (service) IsContentIDMatch(contentID string, digits string) bool {
	return false
}
