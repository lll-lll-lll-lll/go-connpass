package main

import (
	"log"
	"os"

	"github.com/info-api/connpass"
	"github.com/info-api/format"
	"github.com/info-api/markdown"
)

func main() {
	file, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	con, err := connpass.NewConnpass(connpass.USER)
	if err != nil {
		log.Fatal(err)
		return
	}
	initq := map[string]string{"nickname": con.ConnpassUSER}

	con.InitResponse(initq)

	seriesId := con.JoinGroupIdsByComma()
	sm := format.GetForThreeMonthsEvent()
	qd := make(map[string]string)
	qd["series_id"] = seriesId
	qd["count"] = "100"
	qd["ym"] = sm

	con.SetQuery(qd)
	u := con.CreateUrl(con.Query)
	res := con.Request(u)
	defer res.Body.Close()

	err = con.SetResponseBody(res)
	if err != nil {
		log.Fatal(err)
		return
	}
	m := markdown.NewMarkDown()
	m.WriteHandleFunc("", 2, m.WriteTitle)
	m.CreateMd(con.ConnpassResponse)
	s := m.CompleteMdFile(2)
	file.Write([]byte(s))
}
