package connpass

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const CONNPASSAPI string = "https://connpass.com/api/v1/event/?"
const USER string = "Shun_Pei"

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
	u, err := url.Parse(CONNPASSAPI)
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "connpass.com"
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Connpass) InitResponse(query map[string]string) error {
	c.SetQuery(query)
	u := c.CreateUrl(c.Query)
	res := c.Request(u)
	defer res.Body.Close()

	err := c.SetResponseBody(res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
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

// groupidを「,」で繋げる。connpassapiで複数指定は「,」で可能だから
func (c *Connpass) JoinGroupIdsByComma() string {
	seriesId := ""
	gs := c.ConnpassResponse.GetGroups()
	for _, v := range gs {
		v := strconv.Itoa(v)
		seriesId += v + ","
	}
	return seriesId
}
