package connpass

import (
	"fmt"
	"net/url"
)

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

// Query connpass apiにqueryを設定する
// 引数のvaluesが空の場合エラーが発生する
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
