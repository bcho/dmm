package proxy_server

import (
	"context"

	"github.com/bcho/dmm/search"
)

type Options struct {
	Context  context.Context
	HTTPBind string
}

func Main(opt *Options) {
	searchClient := search.NewClient()
	searchService := NewService(WithSearchClient(searchClient))

	serveHTTP(searchService, opt)
}
