package main

import (
	"context"
	"flag"

	"github.com/bcho/dmm/proxy_server"
)

var opt = &proxy_server.Options{
	Context: context.Background(),
}

func main() {
	flag.Parse()

	proxy_server.Main(opt)
}

func init() {
	flag.StringVar(&opt.HTTPBind, "http", ":8080", "http bind address")
}
