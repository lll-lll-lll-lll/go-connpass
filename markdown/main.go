package markdown

import (
	"log"
	"strings"

	"github.com/info-api/connpass"
	"github.com/info-api/format"
)

// mdファイルの全体像を作るメソッド
func (m *MarkDown) CreateMd(response *connpass.ConnpassResponse) *MarkDown {
	for _, v := range response.Events {
		owner := v.Series.Title
		et := v.Title
		eu := v.EventUrl
		es := format.ConvertStartAtTime(v.StartedAt)
		m.WriteTitle(owner, 2)
		m.WriteTitle(et, 3)
		m.WriteHorizon(eu, 1)
		m.WriteHorizon(es, 1)
	}
	return m
}

// MarkDown structに Write(content string, repeat int, method)
type MarkDown struct {
	page     []string
	markdown []*MDElement
}

func (md *MarkDown) NewMDElement() *MDElement {
	mdelement := &MDElement{}
	md.markdown = append(md.markdown, mdelement)
	return mdelement
}

func NewMarkDown() *MarkDown {
	m := new(MarkDown)
	return m
}

// mdのmarkを作成
func (m *MarkDown) CreateMark(mark string, content interface{}, repeat int) string {
	i := interface{}(content)
	n, ok := i.(string)
	if !ok {
		log.Fatal("入力を文字列に変換できませんでした")
	}
	return strings.Repeat(mark, repeat) + " " + n
}

func (m *MarkDown) WriteHorizon(content interface{}, repeat int) {
	markh := "-"
	mark := m.CreateMark(markh, content, repeat)
	m.page = append(m.page, mark)
}

func (m *MarkDown) WriteTitle(content interface{}, repeat int) {
	markt := "#"
	mark := m.CreateMark(markt, content, repeat)
	m.page = append(m.page, mark)
}

// 設定した文字列をつなげて返す
func (m *MarkDown) CompleteMDFile(brNum int) string {
	brs := strings.Repeat("\n", brNum)
	return strings.Join(m.page, brs)
}
