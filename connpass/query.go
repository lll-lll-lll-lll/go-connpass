package connpass

import (
	"log"
	"net/url"
)

func (c *Connpass) SetQuery(values map[string]string) {
	q := url.Values{}
	for k, v := range values {
		q.Add(k, v)
	}
	c.Query = q
}

func (c *Connpass) CreateUrl(q url.Values) string {
	u, err := url.Parse(CONNPASSAPIV1)
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "connpass.com"
	u.RawQuery = q.Encode()
	return u.String()
}
