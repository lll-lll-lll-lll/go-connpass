package connpass

import (
	"log"
	"net/http"
)

func (c *Connpass) Request(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	return res
}
