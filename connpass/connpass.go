package connpass

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	CONNPASSAPIV1 = "https://connpass.com/api/v1/event/?"
)

var DefaultURLValues = url.Values{
	"nickname": []string{""},
}

type Option func(*Client) error

func URL(q url.Values) Option {
	return func(c *Client) error {
		u, err := url.Parse(CONNPASSAPIV1)
		if err != nil {
			return fmt.Errorf("faield to parse connpass api. %w", err)
		}
		u.Scheme = "https"
		u.Host = "connpass.com"
		u.RawQuery = q.Encode()
		c.url = u.String()
		return nil
	}
}

// SetQuery connpass apiにqueryを設定する
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

type Client struct {
	query url.Values
	// connpass apiへのリクエストURLの完成形
	url string
}

func (c *Client) URL() string       { return c.url }
func (c *Client) Query() url.Values { return c.query }
func (c Client) String() string {
	return fmt.Sprintf("query is %v \n url is %s", c.query, c.url)
}

func New(options ...Option) (*Client, error) {
	c := new(Client)
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Client) Do() (*http.Response, error) {
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url %w", err)
	}
	if u.Host != "connpass.com" {
		return nil, fmt.Errorf("host name is not connpass.com")
	}
	res, err := http.Get(c.url)
	if err != nil {
		return nil, fmt.Errorf("failed to do connpass api request. %w", err)
	}
	return res, nil
}

// JoinGroupIDByComma groupidを「,」で繋げる。connpassapiで複数指定は「,」で可能だから
func (r *Response) JoinGroupIDByComma() string {
	var seriesId string
	groupIDs := r.GroupIds()
	for _, v := range groupIDs {
		v := strconv.Itoa(v)
		seriesId += v + ","
	}
	return seriesId
}

// GetGroups 所属してるグループIDを取得
func (r *Response) GroupIds() []int {
	var g = make([]int, len(r.Events))
	for _, v := range r.Events {
		g = append(g, v.Series.Id)
	}
	return g
}

type Event struct {
	Series struct {
		// グループID
		Id int `json:"id"`
		// グループのタイトル
		Title string `json:"title"`
		// グループのconnpass.com 上のURL
		Url string `json:"url"`
	} `json:"series"`
	// 管理者のニックネーム
	OwnerNickname string `json:"owner_nickname"`
	// キャッチコピー
	Catch string `json:"catch"`
	// 概要(HTML形式)
	Description string `json:"description"`
	// connpass.com 上のURL
	EventUrl string `json:"event_url"`
	// Twitterのハッシュタグ
	HashTag string `json:"hash_tag"`
	// イベント開催日時 (ISO-8601形式)
	StartedAt string `json:"started_at"`
	// イベント終了日時 (ISO-8601形式)
	EndedAt string `json:"ended_at"`
	// 管理者の表示名
	OwnerDisplayName string `json:"owner_display_name"`
	// イベント参加タイプ
	EventType string `json:"event_type"`
	// タイトル
	Title string `json:"title"`
	// 開催場所
	Address string `json:"address"`
	// 開催会場
	Place string `json:"place"`
	// 更新日時 (ISO-8601形式)
	UpdatedAt string `json:"updated_at"`
	// イベントID
	EventId int `json:"event_id"`
	// 管理者のID
	OwnerId int `json:"owner_id"`
	// 定員
	Limit int `json:"limit"`
	// 参加者数
	Accepted int `json:"accepted"`
	// 補欠者数
	Waiting int `json:"waiting"`
	// 開催会場の緯度
	Lat string `json:"lat"`
	// 開催会場の経度
	Lon string `json:"lon"`
}

// Response connpass api response
// https://connpass.com/about/api/
type Response struct {
	// レスポンスに含まれる検索結果の件数
	ResultsReturned int `json:"results_returned"`
	// 検索結果の総件数
	ResultsAvailable int `json:"results_available"`
	// 検索の開始位置
	ResultsStart int `json:"results_start"`
	// 検索結果のイベントリスト
	Events []Event `json:"events"`
}
