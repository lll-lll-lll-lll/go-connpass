package connpass

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	CONNPASSAPIV1 = "https://connpass.com/api/v1/event/?"
)

var ErrNotHostNameConnpass = errors.New("host name is not connpass.com")

type RequestURL string

type Client struct {
	UserName string `json:"user_name"` // connpassのユーザ名
	// link: https://connpass.com/about/api/
	//
	// connpass apiから返ってくるレスポンス
	Response *Response  `json:"connpass"`
	Query    url.Values `json:"query"`
	URL      string     `json:"url"` // connpass apiへのリクエストURLの完成形
}

func New() *Client {
	return &Client{}
}

// SetQuery connpass apiにqueryを設定する
// 引数のvaluesが空の場合エラーが発生する
func (c *Client) SetQuery(values map[string]string) error {
	q := url.Values{}
	if len(q) == 0 {
		return errors.New("no query set")
	}
	for k, v := range values {
		q.Add(k, v)
	}
	c.Query = q
	return nil
}

// SetURL CONNPASSAPIV1を解析してc.URLに設定
func (c *Client) SetURL(q url.Values) error {
	u, err := url.Parse(CONNPASSAPIV1)
	if err != nil {
		return fmt.Errorf("connpass apiの解析に失敗しました. %w", err)
	}
	u.Scheme = "https"
	u.Host = "connpass.com"
	u.RawQuery = q.Encode()
	c.URL = u.String()
	return nil
}

// Do connpass apiにリクエストを送る
// URLの解析とホスト名がconnpass.comかどうかチェックしている。
func (c *Client) Do() (*http.Response, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, fmt.Errorf("設定したURLに間違いがあります。%w", err)
	}
	if u.Host != "connpass.com" {
		return nil, ErrNotHostNameConnpass
	}
	res, err := http.Get(c.URL)
	if err != nil {
		return nil, fmt.Errorf("connpass apiへのリクエストに失敗しました。%w", err)
	}
	return res, nil
}

// SetResponse Requestメソッド後のレスポンスをConnpassResponseプロパティにセットする
func (c *Client) SetResponse(res *http.Response) error {
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(body, &c.Response)
	if err != nil {
		return fmt.Errorf("Responseに書き込むのに失敗しました。%w", err)
	}
	return nil
}
