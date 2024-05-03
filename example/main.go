package main

import (
	"log"
	"os"

	"github.com/lll-lll-lll-lll/go-connpass/connpass"
	"github.com/lll-lll-lll-lll/go-connpass/markdown"
)

func main() {
	connpassfunc()
}

func WriteHorizon(m *markdown.MarkDown, content string, repeat int) {
	markh := "-"
	m.AddToPage(markh, content, repeat, 2)
}

func WriteTitle(m *markdown.MarkDown, content string, repeat int) {
	markt := "#"
	m.AddToPage(markt, content, repeat, 2)
}
func WriteBlank(m *markdown.MarkDown, content string, repeat int) {
	mark := "<br>"
	m.AddToPage(mark, content, repeat, 2)
}

// mdファイルの全体像を作るメソッド
func CreateMd(response *connpass.Response, m *markdown.MarkDown) string {
	for _, v := range response.Events {
		owner := v.Series.Title
		et := v.Title
		eu := v.EventUrl
		es := connpass.ConvertStartAtTime(v.StartedAt)
		markt := "#"
		m.AddToPage(markt, owner, 2, 2)
		m.MDHandleFunc(et, 3, WriteTitle)
		m.MDHandleFunc(eu, 1, WriteHorizon)
		m.MDHandleFunc(es, 1, WriteHorizon)
	}
	m.MDHandleFunc("", 1, WriteBlank)
	s := m.CompleteMarkDown(2)
	return s
}

func connpassfunc() {
	file, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	client := connpass.New()
	client.UserName = "Shun_Pei"
	q := map[string]string{"nickname": client.UserName}

	if err := initRequest(client, q); err != nil {
		log.Println(err)
		return
	}

	seriesId := connpass.AggregateGroupIDByComma(client.Response)
	sm := connpass.GetForThreeMonthsEvent()
	qd := make(map[string]string)
	qd["series_id"] = seriesId
	qd["count"] = "100"
	qd["ym"] = sm

	client.SetQuery(qd)
	err = client.SetURL(client.Query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err := client.SetResponse(res); err != nil {
		log.Println(err)
		return
	}

	m := &markdown.MarkDown{}
	s := CreateMd(client.Response, m)
	file.Write([]byte(s))
}

func initRequest(c *connpass.Client, query map[string]string) error {
	c.SetQuery(query)
	err := c.SetURL(c.Query)
	if err != nil {
		return err
	}
	res, _ := c.Do()
	defer res.Body.Close()

	if err := c.SetResponse(res); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
