package proxy_server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeQueryHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request queryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	r.Body.Close()
	return request, nil
}

func decodeIsContentIDMatchHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request isContentIDMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	r.Body.Close()
	return request, nil
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func serveHTTP(s ProxyService, opts *Options) error {
	queryHandler := httptransport.NewServer(
		opts.Context,
		makeQueryEndpoint(s),
		decodeQueryHTTPRequest,
		encodeHTTPResponse,
	)
	isContentIDMatchHandler := httptransport.NewServer(
		opts.Context,
		makeQueryEndpoint(s),
		decodeIsContentIDMatchHTTPRequest,
		encodeHTTPResponse,
	)

	http.Handle("/api/v1/query", queryHandler)
	http.Handle("/api/v1/content_id.match", isContentIDMatchHandler)

	log.Printf("http service started at %s", opts.HTTPBind)
	return http.ListenAndServe(opts.HTTPBind, nil)
}
