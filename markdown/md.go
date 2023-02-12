package markdown

import (
	"fmt"
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
func (m *MarkDown) CreateMark(mark string, content interface{}, repeat int) (string, error) {
	i := interface{}(content)
	n, ok := i.(string)
	if !ok {
		return "", fmt.Errorf("入力を文字列に変換するのに失敗しました")
	}
	return strings.Repeat(mark, repeat) + " " + n, nil
}

func (m *MarkDown) AddToPage(mark string, content interface{}, repeat int) error {
	melement, err := m.CreateMark(mark, content, repeat)
	if err != nil {
		return err
	}
	m.page = append(m.page, melement)
	return nil
}

// 設定した文字列をつなげて返す
func (m *MarkDown) CompleteMarkDown(branknum int) string {
	brs := strings.Repeat("\n", branknum)
	return strings.Join(m.page, brs)
}
