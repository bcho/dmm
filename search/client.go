package search

import (
	"fmt"
	"net/http"
)

type clientOpt func(*Client)

// WithQueryURLBuilder binds custom query url builder to client.
func WithQueryURLBuilder(b func(string) string) clientOpt {
	return func(c *Client) {
		c.queryURLBuilder = b
	}
}

// WithHTTPClient binds a http client to client.
func WithHTTPClient(h *http.Client) clientOpt {
	return func(c *Client) {
		c.httpClient = h
	}
}

// Client represents a search client.
type Client struct {
	queryURLBuilder func(string) string
	httpClient      *http.Client
}

// NewClient creates a search client.
func NewClient(opts ...clientOpt) *Client {
	c := &Client{}

	for _, o := range opts {
		o(c)
	}

	if c.queryURLBuilder == nil {
		c.queryURLBuilder = func(s string) string {
			return fmt.Sprintf("http://www.dmm.co.jp/search/=/searchstr=%s", s)
		}
	}

	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	return c
}

var defaultClient = NewClient()
