package search

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c.queryURLBuilder == nil {
		t.Fatalf("queryURLBuilder should not be nil")
	}
	if c.httpClient == nil {
		t.Fatalf("httpClient should not be nil")
	}
}
