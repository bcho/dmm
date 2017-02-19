package main

import (
	"flag"
	"log"

	"github.com/bcho/dmm/search"
)

func main() {
	q := flag.String("q", "", "")

	flag.Parse()

	products, err := search.Query(&search.QueryOpts{
		Keyword: *q,
	})

	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, prod := range products {
		log.Printf("%s", prod)
	}
}
