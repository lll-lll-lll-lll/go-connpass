package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lll-lll-lll-lll/go-connpass/connpass"
	"github.com/lll-lll-lll-lll/go-connpass/markdown"
)

func main() {
	connpassfunc()
}

func WriteHorizon(m *markdown.MarkDown, content string, repeat int) {
	markh := "-"
	markElem := m.Mark(markh, repeat)
	m.Add(markElem + " " + content)
	m.AddBr(2)
}

func WriteTitle(m *markdown.MarkDown, content string, repeat int) {
	markt := "#"
	markElem := m.Mark(markt, repeat)
	m.Add(markElem + " " + content)
	m.AddBr(2)
}
func WriteBlank(m *markdown.MarkDown, content string, repeat int) {
	mark := "<br>"
	markElem := m.Mark(mark, repeat)
	m.Add(markElem + " " + content)
	m.AddBr(2)
}

// mdファイルの全体像を作るメソッド
func CreateMd(response *connpass.Response, m *markdown.MarkDown) string {
	for _, v := range response.Events {
		owner := v.Series.Title
		et := v.Title
		eu := v.EventUrl
		es := convertStartAtTime(v.StartedAt)
		WriteTitle(m, owner, 3)
		WriteTitle(m, et, 3)
		WriteHorizon(m, eu, 1)
		WriteHorizon(m, es, 1)
	}
	return m.String()
}

func connpassfunc() {
	file, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	client := &connpass.Client{}
	initRes, err := initRequest(client)
	if err != nil {
		log.Fatal(err)
	}
	groupIDs := initRes.GroupIds()
	seriesId := initRes.JoinGroupIDs(groupIDs)
	sm := getForThreeMonthsEvent()
	qd := make(map[string]string)
	qd["series_id"] = seriesId
	qd["count"] = "100"
	qd["ym"] = sm

	res, err := client.Do(context.Background(), connpass.Query(qd), connpass.URLV1())
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var cr connpass.Response
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &cr); err != nil {
		log.Fatal(err)
	}

	m := &markdown.MarkDown{}
	s := CreateMd(&cr, m)
	file.Write([]byte(s))
}

func initRequest(c *connpass.Client) (*connpass.Response, error) {
	q := map[string]string{"nickname": "Shun_Pei"}
	res, _ := c.Do(context.Background(),connpass.Query(q), connpass.URLV1())
	defer res.Body.Close()

	var cRes connpass.Response
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &cRes); err != nil {
		return nil, fmt.Errorf("Responseに書き込むのに失敗しました。%w", err)
	}

	return &cRes, nil
}

// 今月を含めた３月分のイベントを取得
func getForThreeMonthsEvent() string {
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
	sm, tm := checkMonthFormat(nm)
	// a := yearmonthsja.Replace(fmt.Sprintf("%s", tm.String()))

	f := yearmonthsja.Replace(fmt.Sprintf("%d%s", now.Year(), nm.String()))
	s := yearmonthsja.Replace(fmt.Sprintf("%d%s", now.Year(), sm))
	t := yearmonthsja.Replace(fmt.Sprintf("%d%s", now.Year(), tm))
	return f + "," + s + "," + t
}

// 12で割るときに12月だけ0が返ってくるので、その時だけ"12"文字列を返す
func checkMonthFormat(nm time.Month) (string, string) {
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
func convertStartAtTime(startedAt string) string {
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
