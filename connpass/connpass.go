package connpass

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	CONNPASSAPIV1 = "https://connpass.com/api/v1/event/?"
)

var ErrNotHostNameConnpass = errors.New("host name is not connpass.com")

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
	if len(values) == 0 {
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

// AggregateGroupIDByComma groupidを「,」で繋げる。connpassapiで複数指定は「,」で可能だから
func AggregateGroupIDByComma(res *Response) string {
	var seriesId string
	groupIDs := res.GetGroupIds()
	for _, v := range groupIDs {
		v := strconv.Itoa(v)
		seriesId += v + ","
	}
	return seriesId
}

// GetGroups 所属してるグループIDを取得
func (c *Response) GetGroupIds() []int {
	var g []int
	for _, v := range c.Events {
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

// Response コンパスapiのレスを持つ
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
