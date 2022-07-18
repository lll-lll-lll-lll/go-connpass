package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const CONNPASSAPI string = "https://connpass.com/api/v1/event/?"
const USER string = "Shun_Pei"

// group idを取得するのが目的
func initRequestConnpass() *ConnpassResponse {
	qm := map[string]string{"nickname": USER}
	q := CreateQuery(qm)
	u := CreateUrl(q)
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	response := NewConnpassResponse()
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}
	return response
}

func main() {
	file, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	initres := initRequestConnpass()

	qd := make(map[string]string)
	// 所属してるグループId取得
	gs := initres.GetGroups()
	// groupidを「,」で繋げる。connpassapiで複数指定は「,」で可能だから
	seriesId := ""
	for _, v := range gs {
		v := strconv.Itoa(v)
		seriesId += v + ","
	}
	sm := GetForThreeMonthsEvent()
	qd["series_id"] = seriesId
	qd["count"] = "100"
	qd["ym"] = sm

	q := CreateQuery(qd)
	u := CreateUrl(q)
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	response := NewConnpassResponse()
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	mn := new(MarkDown)
	m := CreateMd(mn, response)
	s := m.CompleteMdFile(2)
	file.Write([]byte(s))

}

// 今月を含めた３月分のイベントを取得
func GetForThreeMonthsEvent() string {
	now := time.Now()
	yearmonthsja := strings.NewReplacer(
		"January", "01",
		"February", "02",
		"March", "03",
		"April", "04",
		"May", "05",
		"June", "06",
		"July", "07",
		"August", "08",
		"September", "09",
		"October", "10",
		"November", "11",
		"December", "12",
	)
	// 13月をなくすために12で割った余を入れる
	nm := now.Month()
	sm, tm := CheckMonthFormat(nm)
	// a := yearmonthsja.Replace(fmt.Sprintf("%s", tm.String()))

	f := yearmonthsja.Replace(fmt.Sprintf("%d%s", now.Year(), nm.String()))
	s := yearmonthsja.Replace(fmt.Sprintf("%d%s", now.Year(), sm))
	t := yearmonthsja.Replace(fmt.Sprintf("%d%s", now.Year(), tm))
	return f + "," + s + "," + t
}

// 12で割るときに12月だけ0が返ってくるので、その時だけ"12"文字列を返す
func CheckMonthFormat(nm time.Month) (string, string) {
	ze := "%!Month(0)"
	sm := (nm + 1) % 12
	tm := (nm + 2) % 12
	if sm.String() == ze {
		return "12", tm.String()
	} else if tm.String() == ze {
		return sm.String(), "12"
	}
	return sm.String(), tm.String()
}

// 時刻を見やすいように変更
func ConvertStartAtTime(startedAt string) string {
	weekdaymonthja := strings.NewReplacer(
		"Sunday", "日",
		"Monday", "月",
		"Tueday", "火",
		"Wednesday", "水",
		"Thursday", "木",
		"Friday", "金",
		"Saturday", "土",
		"January", "1月",
		"February", "2月",
		"March", "3月",
		"April", "4月",
		"May", "5月",
		"June", "6月",
		"July", "7月",
		"August", "8月",
		"September", "9月",
		"October", "10月",
		"November", "11月",
		"December", "12月",
	)
	p, err := time.Parse(time.RFC3339, startedAt)
	if err != nil {
		log.Fatal(err)
	}
	f := func(p time.Time) string {
		if p.Minute() == 0 {
			return "00"
		} else {
			return strconv.Itoa(p.Minute())
		}
	}
	str := fmt.Sprintf("%s%d日(%s) %d:%s ~", p.Month().String(), p.Day(), p.Weekday(), p.Hour(), f(p))
	return weekdaymonthja.Replace(str)
}

// mdファイルの全体像を作るメソッド
func CreateMd(m *MarkDown, response *ConnpassResponse) *MarkDown {
	for _, v := range response.Events {
		owner := v.Series.Title
		et := v.Title
		eu := v.EventUrl
		es := ConvertStartAtTime(v.StartedAt)
		m.WriteTitle(owner, 2)
		m.WriteTitle(et, 3)
		m.WriteHorizon(eu, 1)
		m.WriteHorizon(es, 1)
	}
	return m
}

// クエリを作成
func CreateQuery(values map[string]string) url.Values {
	q := url.Values{}
	for k, v := range values {
		q.Add(k, v)
	}
	return q
}

// CreateUrl Urlにクエリーを設定してurlを返す
func CreateUrl(q url.Values) string {
	u, err := url.Parse(CONNPASSAPI)
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "connpass.com"
	u.RawQuery = q.Encode()
	return u.String()
}

type MarkDown struct {
	page []string
}

func NewMarkDown() *MarkDown {
	m := new(MarkDown)
	return m
}

// mdのmarkを作成
func (m *MarkDown) CreateMark(mark string, content string, repeat int) string {
	return strings.Repeat(mark, repeat) + " " + content
}

func (m *MarkDown) WriteHorizon(content string, repeat int) *MarkDown {
	markh := "-"
	mark := m.CreateMark(markh, content, repeat)
	m.page = append(m.page, mark)
	return m
}

func (m *MarkDown) WriteTitle(content string, repeat int) *MarkDown {
	markt := "#"
	mark := m.CreateMark(markt, content, repeat)
	m.page = append(m.page, mark)
	return m
}

// 設定した文字列をつなげて返す
func (m *MarkDown) CompleteMdFile(brNum int) string {
	brs := strings.Repeat("\n", brNum)
	return strings.Join(m.page, brs)
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

// ConnpassResponse コンパスapiのレスを持つ
type ConnpassResponse struct {
	// 含まれる検索結果の件数
	ResultsReturned int `json:"results_returned"`
	// 検索結果の総件数
	ResultsAvailable int `json:"results_available"`
	// 検索の開始位置
	ResultsStart int `json:"results_start"`
	// 検索結果のイベントリスト
	Events []Event `json:"events"`
}

func NewConnpassResponse() *ConnpassResponse {
	return &ConnpassResponse{}
}

func (c *ConnpassResponse) GetGroups() []int {
	var g []int
	for _, v := range c.Events {
		g = append(g, v.Series.Id)
	}
	return g
}
