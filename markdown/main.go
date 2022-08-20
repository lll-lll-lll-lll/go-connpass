package markdown

import (
	"log"
	"strings"
)

// MarkDown structに Write(content string, repeat int, method)
type MarkDown struct {
	page []string
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

func (m *MarkDown) AddToPage(mark string, content interface{}, repeat int) {
	melement := m.CreateMark(mark, content, repeat)
	m.page = append(m.page, melement)
}

func (m *MarkDown) WriteHorizon(content interface{}, repeat int) {
	markh := "-"
	m.AddToPage(markh, content, repeat)
}

func (m *MarkDown) WriteTitle(content interface{}, repeat int) {
	markt := "#"
	m.AddToPage(markt, content, repeat)
}

// 設定した文字列をつなげて返す
func (m *MarkDown) CompleteMDFile(branknum int) string {
	brs := strings.Repeat("\n", branknum)
	return strings.Join(m.page, brs)
}
