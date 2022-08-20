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
