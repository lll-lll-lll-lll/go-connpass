package markdown

import (
	"strings"
)

// MarkDown structに Write(content string, repeat int, method)
type MarkDown struct {
	s *strings.Builder
}

func (m MarkDown) String() string {
	return m.s.String()
}

// mdのmarkを作成
func (m *MarkDown) generateMark(mark string, content string, repeat int) string {
	return strings.Repeat(mark, repeat) + " " + content
}

func (m *MarkDown) AddToPage(mark, content string, repeat, branknum int) error {
	brs := strings.Repeat("\n", branknum)
	melement := m.generateMark(mark, content, repeat)
	if _, err := m.s.WriteString(melement + brs); err != nil {
		return err
	}
	return nil
}

// // 設定した文字列をつなげて返す
// func (m *MarkDown) CompleteMarkDown(branknum int) string {
// 	brs := strings.Repeat("\n", branknum)
// 	return strings.Join(m.page, brs)
// }
