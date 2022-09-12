package main

import (
	"log"
	"os"

	"github.com/lll-lll-lll-lll/conread/connpass"
	"github.com/lll-lll-lll-lll/conread/format"
	"github.com/lll-lll-lll-lll/conread/markdown"
)

func main() {
	connpassfunc()
}

func WriteHorizon(m *markdown.MarkDown, content interface{}, repeat int) {
	markh := "-"
	m.AddToPage(markh, content, repeat)
}

func WriteTitle(m *markdown.MarkDown, content interface{}, repeat int) {
	markt := "#"
	m.AddToPage(markt, content, repeat)
}
func WriteBlank(m *markdown.MarkDown, content interface{}, repeat int) {
	mark := "<br>"
	m.AddToPage(mark, content, repeat)
}

// mdファイルの全体像を作るメソッド
func CreateMd(response *connpass.ConnpassResponse, m *markdown.MarkDown) string {
	for _, v := range response.Events {
		owner := v.Series.Title
		et := v.Title
		eu := v.EventUrl
		es := format.ConvertStartAtTime(v.StartedAt)
		m.MDHandleFunc(owner, 2, WriteTitle)
		m.MDHandleFunc(et, 3, WriteTitle)
		m.MDHandleFunc(eu, 1, WriteHorizon)
		m.MDHandleFunc(es, 1, WriteHorizon)
	}
	m.MDHandleFunc("", 1, WriteBlank)
	s := m.CompleteMDFile(2)
	return s
}

func connpassfunc() {
	file, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	con := connpass.NewConnpass()
	con.ConnpassUSER = "Shun_Pei"
	q := map[string]string{"nickname": con.ConnpassUSER}

	if err := initRequest(con, q); err != nil {
		log.Println(err)
		return
	}

	seriesId := con.JoinGroupIdsByComma()
	sm := format.GetForThreeMonthsEvent()
	qd := make(map[string]string)
	qd["series_id"] = seriesId
	qd["count"] = "100"
	qd["ym"] = sm

	createdQuery := connpass.CreateQuery(qd)
	con.Query = createdQuery
	u := con.CreateUrl(con.Query)
	res := con.Request(u)
	defer res.Body.Close()

	err = con.SetResponse(res)
	if err != nil {
		log.Fatal(err)
		return
	}

	m := markdown.NewMarkDown()
	s := CreateMd(con.ConnpassResponse, m)
	file.Write([]byte(s))
}

func defaultfunc() {
	file, err := os.Create("default.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Remove("default.md")
	m := markdown.NewMarkDown()
	m.MDHandleFunc("Test Write Title", 2, WriteTitle)
	m.MDHandleFunc("Test Write Horizon", 3, WriteHorizon)
	s := m.CompleteMDFile(2)
	file.Write([]byte(s))
}

func initRequest(c *connpass.Connpass, query map[string]string) error {
	q := connpass.CreateQuery(query)
	c.Query = q
	u := c.CreateUrl(c.Query)
	res := c.Request(u)
	defer res.Body.Close()

	if err := c.SetResponse(res); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
