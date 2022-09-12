package connpass

import (
	"log"
	"net/http"
)

func (c *Connpass) InitRequest(query map[string]string) error {
	c.SetQuery(query)
	u := c.CreateUrl(c.Query)
	res := c.Request(u)
	defer res.Body.Close()

	err := c.SetResponse(res)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *Connpass) Request(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	return res
}
