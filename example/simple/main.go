package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/lll-lll-lll-lll/go-connpass/connpass"
)

func main() {
	client := &connpass.Client{}
	req := &connpass.EventRequest{}
	req.SetURL(connpass.CONNPASSAPI_EVENT_V1 + "?")
	req.NickName = []string{"your connpass nickname"}
	res, _ := client.Do(context.Background(), req)
	defer res.Body.Close()

	var cRes connpass.EventResponse
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &cRes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cRes)
}
