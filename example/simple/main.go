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
	q := map[string]string{"nickname": "your connpass user name"}
	res, _ := client.Do(context.Background(), connpass.URL(q))
	defer res.Body.Close()

	var cRes connpass.Response
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &cRes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cRes)
}
