package connpass

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

func (c *Connpass) InitRequest(query map[string]string) error {
	c.SetQuery(query)
	u := c.CreateUrl(c.Query)
	res := c.Request(u)
	defer res.Body.Close()

	err := c.SetResponse(res)
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

//SetResponse Requestメソッド後のレスポンスをConnpassResponseプロパティにセットする
func (c *Connpass) SetResponse(res *http.Response) error {
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(body, &c.ConnpassResponse)
	if err != nil {
		return err
	}
	return nil
}

//JoinGroupIdsByComma groupidを「,」で繋げる。connpassapiで複数指定は「,」で可能だから
func (c *Connpass) JoinGroupIdsByComma() string {
	var seriesId string
	gs := c.ConnpassResponse.GetGroupIds()
	for _, v := range gs {
		v := strconv.Itoa(v)
		seriesId += v + ","
	}
	return seriesId
}
