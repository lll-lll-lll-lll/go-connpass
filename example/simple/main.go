package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/lll-lll-lll-lll/go-connpass/connpass"
)

func main() {
	client := &connpass.Client{}
	req := &connpass.EventRequest{}
	req.Path = connpass.EVENT_PATH
	req.NickName = []string{"your connpass nickname"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, _ := client.Do(ctx, req)
	defer res.Body.Close()

	var cRes connpass.EventResponse
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &cRes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cRes)
}
