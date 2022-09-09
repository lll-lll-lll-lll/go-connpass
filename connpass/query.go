package connpass

import (
	"log"
	"net/url"
)

const CONNPASSAPIV1 = "https://connpass.com/api/v1/event/?"

type Connpass struct {
	ConnpassUSER     string            `json:"user"`
	ConnpassResponse *ConnpassResponse `json:"connpass"`
	Query            url.Values        `json:"query"`
}

func NewConnpass(user string) (*Connpass, error) {
	c := &Connpass{
		ConnpassUSER: user,
	}
	return c, nil
}

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
