package proxy_server

import "github.com/bcho/dmm"

// ProxyService implements a cached query proxy server.
type ProxyService interface {
	Query(string) ([]dmm.Product, error)
	IsContentIDMatch(string, string) bool
}
