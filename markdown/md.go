package markdown

import (
	"strings"
)

// MarkDown structに Write(content string, repeat int, method)
type MarkDown struct {
	s strings.Builder
}

func (m MarkDown) String() string {
	return m.s.String()
}

func (m *MarkDown) Mark(mark string, repeat int) string {
	return strings.Repeat(mark, repeat)
}

func (m *MarkDown) AddBr(branknum int) (int, error) {
	brs := strings.Repeat("\n", branknum)
	return m.s.WriteString(brs)
}

func (m *MarkDown) Add(content string) (int, error) {
	return m.s.WriteString(content)
}
