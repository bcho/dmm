package proxy_server

import (
	"fmt"
	"time"

	"github.com/bcho/dmm"
	"github.com/bcho/dmm/search"
	cache "github.com/patrickmn/go-cache"
)

type serviceOpt func(s *service)

func WithSearchClient(client *search.Client) serviceOpt {
	return func(s *service) {
		s.client = client
	}
}

type service struct {
	client *search.Client
	cache  *cache.Cache
}

func NewService(opts ...serviceOpt) *service {
	s := &service{
		cache: cache.New(24*time.Hour, 1*time.Hour),
	}

	for _, o := range opts {
		o(s)
	}

	if s.client == nil {
		panic("search client is required")
	}

	return s
}

func (s service) Query(keyword string) (rv []dmm.Product, err error) {
	rv, err = s.queryFromCache(keyword)
	if err != nil {
		return
	}

	if rv != nil {
		// hit from cache
		return
	}

	rv, err = s.client.Query(&search.QueryOpts{keyword})
	if rv != nil {
		// hit from remote
		err = s.cacheResult(keyword, rv)
	}

	return
}

func (service) IsContentIDMatch(contentID string, digits string) bool {
	return false
}

func (s service) queryCacheKey(keyword string) string {
	return fmt.Sprintf("query-%s", keyword)
}

func (s service) queryFromCache(keyword string) (rv []dmm.Product, err error) {
	r, found := s.cache.Get(s.queryCacheKey(keyword))
	if !found {
		return nil, nil
	}

	return r.([]dmm.Product), nil
}

func (s service) cacheResult(keyword string, rv []dmm.Product) error {
	s.cache.Set(s.queryCacheKey(keyword), rv, cache.DefaultExpiration)
	return nil
}
