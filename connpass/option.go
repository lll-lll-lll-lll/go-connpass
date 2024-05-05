package connpass

import (
	"fmt"
	"net/url"
)

var DefaultURLValues = url.Values{
	"nickname":  []string{""},
	"count":     []string{"100"},
	"ym":        []string{"202405"},
	"series_id": []string{"3494,13377,6108,"},
}

type Option func(*Client) error

func URLV1() Option {
	return func(c *Client) error {
		u, err := url.Parse(CONNPASSAPIV1)
		if err != nil {
			return fmt.Errorf("faield to parse connpass api. %w", err)
		}
		u.Scheme = "https"
		u.Host = "connpass.com"
		u.RawQuery = c.query.Encode()
		c.url = u.String()
		return nil
	}
}

func Query(values map[string]string) Option {
	return func(c *Client) error {
		q := url.Values{}
		if len(values) == 0 {
			c.query = DefaultURLValues
			return nil
		}
		for k, v := range values {
			q.Add(k, v)
		}
		c.query = q
		return nil
	}
}
