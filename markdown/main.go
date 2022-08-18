package markdown

import (
	"strings"

	"github.com/info-api/connpass"
)

// mdファイルの全体像を作るメソッド
func (m *MarkDown) CreateMd(response *connpass.ConnpassResponse) *MarkDown {
	for _, v := range response.Events {
		owner := v.Series.Title
		et := v.Title
		eu := v.EventUrl
		es := ConvertStartAtTime(v.StartedAt)
		m.WriteTitle(owner, 2)
		m.WriteTitle(et, 3)
		m.WriteHorizon(eu, 1)
		m.WriteHorizon(es, 1)
	}
	return m
}

type MarkDown struct {
	page []string
}

func NewMarkDown() *MarkDown {
	m := new(MarkDown)
	return m
}

// mdのmarkを作成
func (m *MarkDown) CreateMark(mark string, content string, repeat int) string {
	return strings.Repeat(mark, repeat) + " " + content
}

func (m *MarkDown) WriteHorizon(content string, repeat int) *MarkDown {
	markh := "-"
	mark := m.CreateMark(markh, content, repeat)
	m.page = append(m.page, mark)
	return m
}

func (m *MarkDown) WriteTitle(content string, repeat int) *MarkDown {
	markt := "#"
	mark := m.CreateMark(markt, content, repeat)
	m.page = append(m.page, mark)
	return m
}

// 設定した文字列をつなげて返す
func (m *MarkDown) CompleteMdFile(brNum int) string {
	brs := strings.Repeat("\n", brNum)
	return strings.Join(m.page, brs)
}
