package connpass

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	CONNPASSAPI_V1 string = "https://connpass.com/api/v1"
)

type Client struct{}

func (c *Client) Do(ctx context.Context, req ConnpassRequest) (*http.Response, error) {
	if req.URL() == "" {
		return nil, fmt.Errorf("request url is empty")
	}
	q := req.ToQueryParameter()
	u, err := url.Parse(req.URL())
	if err != nil {
		return nil, fmt.Errorf("faield to parse connpass api. %w", err)
	}
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to do connpass api request. %w", err)
	}
	return res, nil
}
