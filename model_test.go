package dmm

import (
	"log"
	"testing"
)

func TestParseContentID(t *testing.T) {
	var cs = []struct {
		input, expected string
	}{
		{
			"http://foobar/=/cid=12345",
			"12345",
		},
		{
			"http://foobar/=/cid=foobar12345/?a=b",
			"foobar12345",
		},
		{
			"http://foobar",
			"",
		},
	}

	for _, c := range cs {
		if ParseContentID(c.input) != c.expected {
			log.Printf("parse %s failed", c.input)
		}
	}
}

func TestParseLargeCoverURL(t *testing.T) {
	var cs = []struct {
		input, expected string
	}{
		{
			"//pic.foobar/apt.jpg",
			"http://pic.foobar/apl.jpg",
		},
		{
			"//pic.foobar/aps.jpg",
			"http://pic.foobar/apl.jpg",
		},
		{
			"//pic.foobar/apl.jpg",
			"http://pic.foobar/apl.jpg",
		},
		{
			"https://foobar/apt.jpg",
			"https://foobar/apl.jpg",
		},
	}
	for _, c := range cs {
		if ParseLargeCoverURL(c.input) != c.expected {
			log.Printf("parse %s failed", c.input)
		}
	}
}
