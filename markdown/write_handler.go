package markdown

// net/httpを参考にした. https://pkg.go.dev/net/http

type Write interface{}

type WriteHandlerFunc func(content string, repeat int)

func (w WriteHandlerFunc) WriteFunc(content string, repeat int) {
	w(content, repeat)
}

type WriteHandler interface {
	WriteFunc(content string, repeat int)
}

func (md *MarkDown) WriteHandle(content string, repeat int, write WriteHandler) {}

func (md *MarkDown) WriteHandleFunc(content string, repeat int, write func(content string, repeat int)) {
	md.WriteHandle(content, repeat, WriteHandlerFunc(write))
}
