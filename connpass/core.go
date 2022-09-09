package connpass

import "net/url"

const CONNPASSAPIV1 = "https://connpass.com/api/v1/event/?"

type Connpass struct {
	ConnpassUSER     string            `json:"user"`
	ConnpassResponse *ConnpassResponse `json:"connpass"`
	Query            url.Values        `json:"query"`
}

func NewConnpass(user string) (*Connpass, error) {
	c := &Connpass{
		ConnpassUSER: user,
	}
	return c, nil
}
