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
	connpass, err := connpass.NewConnpass(connpass.USER)
	if err != nil {
		log.Fatal(err)
		return
	}
	initq := map[string]string{"nickname": connpass.ConnpassUSER}

	connpass.InitResponse(initq)

	seriesId := connpass.JoinGroupIdsByComma()
	sm := format.GetForThreeMonthsEvent()
	qd := make(map[string]string)
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
	m := markdown.NewMarkDown()
	m.CreateMd(connpass.ConnpassResponse)
	s := m.CompleteMdFile(2)
	file.Write([]byte(s))
}
