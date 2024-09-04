package connpass

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	CONNPASSAPI_EVENT_V1 string = "https://connpass.com/api/v1/event/"
	CONNPASSAPI_USER_V1  string = "https://connpass.com/api/v1/user/"
)

type Client struct{}

func (c *Client) Do(ctx context.Context, req ConnpassRequest) (*http.Response, error) {
	q := req.ToURLVal()
	u, err := url.Parse(req.URL())
	if err != nil {
		return nil, fmt.Errorf("faield to parse connpass api. %w", err)
	}
	u.RawQuery = q.Encode()
	if u.Scheme != "https" || u.Hostname() != "connpass.com" {
		return nil, fmt.Errorf("host name is not connpass.com")
	}
	res, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to do connpass api request. %w", err)
	}
	return res, nil
}
