package proxy_server

import (
	"context"

	"github.com/bcho/dmm"
	"github.com/go-kit/kit/endpoint"
)

type queryRequest struct {
	Keyword string `json:"keyword"`
}
type queryResponse struct {
	Products []dmm.Product `json:"products"`
	Err      string        `json:"err,omitempty"`
}

type isContentIDMatchRequest struct {
	ContentID string `json:"content_id"`
	Digits    string `json:"digits"`
}

type isContentIDMatchResponse struct {
	Matched bool `json:"matched"`
}

func makeQueryEndpoint(s ProxyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(queryRequest)
		products, err := s.Query(req.Keyword)
		if err != nil {
			return queryResponse{nil, err.Error()}, nil
		}
		return queryResponse{products, ""}, nil
	}
}

func makeIsContentIDMatchEndpoint(s ProxyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(isContentIDMatchRequest)
		isMatched := s.IsContentIDMatch(req.ContentID, req.Digits)
		return isContentIDMatchResponse{isMatched}, nil
	}
}
