package connpass

import (
	"fmt"
	"net/url"
)

type Option func(*Client) error

func URL(queryKeyVal map[string]string) Option {
	return func(c *Client) error {
		q := url.Values{}
		if len(queryKeyVal) == 0 {
			return fmt.Errorf("query key value not set")
		}
		for k, v := range queryKeyVal {
			q.Add(k, v)
		}
		c.query = q
		u, err := url.Parse(CONNPASSAPIV1)
		if err != nil {
			return fmt.Errorf("faield to parse connpass api. %w", err)
		}
		u.RawQuery = q.Encode()
		c.url = u.String()
		return nil
	}
}
