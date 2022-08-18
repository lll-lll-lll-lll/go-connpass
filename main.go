package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/info-api/connpass"
)

func main() {
	file, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	connpass, err := connpass.NewConnpass(connpass.USER)
	if err != nil {
		log.Fatal(err)
		return
	}

	qd := make(map[string]string)
	seriesId := connpass.JoinGroupIdsByComma()
	sm := GetForThreeMonthsEvent()
	qd["series_id"] = seriesId
	qd["count"] = "100"
	qd["ym"] = sm

	connpass.SetQuery(qd)
	u := connpass.CreateUrl(connpass.Query)
	res := connpass.Request(u)
	defer res.Body.Close()

	err = connpass.SetResponseBody(res)
	if err != nil {
		log.Fatal(err)
		return
	}

	mn := new(MarkDown)
	m := CreateMd(mn, connpass.ConnpassResponse)
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
func CreateMd(m *MarkDown, response *connpass.ConnpassResponse) *MarkDown {
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

// CreateUrl Urlにクエリーを設定してurlを返す
func CreateUrl(q url.Values) string {
	u, err := url.Parse(connpass.CONNPASSAPI)
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
