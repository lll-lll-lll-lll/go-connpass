package connpass

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	CONNPASSAPIV1 = "https://connpass.com/api/v1/event/?"
)

var DefaultURLValues = url.Values{
	"nickname": []string{""},
}

type Client struct {
	query url.Values
	url   string
}

func New(options ...Option) (*Client, error) {
	c := new(Client)
	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Client) Do(options ...Option) (*http.Response, error) {
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
