package markdown

// net/httpを参考にした. https://pkg.go.dev/net/http

type MDHandlerFunc func(md *MarkDown, content string, repeat int)

func (w MDHandlerFunc) MDFunc(md *MarkDown, content string, repeat int) {
	w(md, content, repeat)
}

type MDHandler interface {
	MDFunc(md *MarkDown, content string, repeat int)
}

func (md *MarkDown) MDHandle(content string, repeat int, write MDHandler) {
	write.MDFunc(md, content, repeat)
}

func (md *MarkDown) MDHandleFunc(content string, repeat int, write func(md *MarkDown, content string, repeat int)) {
	md.MDHandle(content, repeat, MDHandlerFunc(write))
}
