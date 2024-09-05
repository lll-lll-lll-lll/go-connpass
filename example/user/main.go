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
	connpassClient := &connpass.Client{}
	req := &connpass.UserRequest{}
	req.SetURL(connpass.CONNPASSAPI_USER_V1 + "?")
	q := map[string][]string{"nickname": {"Shun_Pei"}}
	req.SetURLQuery(q)
	res, err := connpassClient.Do(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var cRes connpass.UserResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &cRes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cRes)

}
