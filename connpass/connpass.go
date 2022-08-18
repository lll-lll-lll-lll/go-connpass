package connpass

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

const CONNPASSAPI string = "https://connpass.com/api/v1/event/?"
const USER string = "Shun_Pei"

type Connpass struct {
	ConnpassUSER     string            `json:"user"`
	ConnpassResponse *ConnpassResponse `json:"connpass"`
	Query            url.Values        `json:"query"`
}

func NewConnpass(user string) (*Connpass, error) {
	var err error
	c := &Connpass{
		ConnpassUSER: user,
	}
	c.ConnpassResponse, err = c.InitResponse()
	if err != nil {
		return nil, err
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
	u, err := url.Parse(CONNPASSAPI)
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "connpass.com"
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Connpass) InitResponse() (*ConnpassResponse, error) {
	qm := map[string]string{"nickname": c.ConnpassUSER}
	c.SetQuery(qm)
	u := c.CreateUrl(c.Query)
	res := c.Request(u)
	defer res.Body.Close()

	err := c.SetResponseBody(res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return c.ConnpassResponse, nil
}

func (c *Connpass) Request(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (c *Connpass) SetResponseBody(res *http.Response) error {
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(body, &c.ConnpassResponse)
	if err != nil {
		return err
	}
	return nil
}
