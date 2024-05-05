package connpass

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	CONNPASSAPIV1 = "https://connpass.com/api/v1/event/?"
)

type Client struct {
	query url.Values
	url   string
}

func (c *Client) Do(ctx context.Context, options ...Option) (*http.Response, error) {
	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url %w", err)
	}
	if u.Scheme != "https" || u.Hostname() != "connpass.com" {
		return nil, fmt.Errorf("host name is not connpass.com")
	}
	res, err := http.Get(c.url)
	if err != nil {
		return nil, fmt.Errorf("failed to do connpass api request. %w", err)
	}
	return res, nil
}

func (c *Client) URL() string { return c.url }

func (c *Client) Query() url.Values { return c.query }
