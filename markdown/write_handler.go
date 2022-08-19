package markdown

// net/httpを参考にした. https://pkg.go.dev/net/http

type Write interface{}

type WriteHandlerFunc func(content interface{}, repeat int, write *Write)

type WriteHandler interface {
	WriteFunc(content interface{}, repeat int, write *Write)
}

func (md *MarkDown) WriteHandle(content interface{}, repeat int, write WriteHandler) {}

func (md *MarkDown) WriteHandleFunc(content interface{}, repeat int, write func(content interface{}, repeat int, write *Write)) {
	md.WriteHandle(content, repeat, WriteHandlerFunc(write))
}

func (w WriteHandlerFunc) WriteFunc(content interface{}, repeat int, write *Write) {
	w(content, repeat, write)
}
