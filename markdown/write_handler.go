package markdown

// net/httpを参考にした. https://pkg.go.dev/net/http

type WriteHandlerFunc func(md *MarkDown, content interface{}, repeat int)

func (w WriteHandlerFunc) WriteFunc(md *MarkDown, content interface{}, repeat int) {
	w(md, content, repeat)
}

type WriteHandler interface {
	WriteFunc(md *MarkDown, content interface{}, repeat int)
}

func (md *MarkDown) WriteHandle(content interface{}, repeat int, write WriteHandler) {
	write.WriteFunc(md, content, repeat)
}

func (md *MarkDown) WriteHandleFunc(content interface{}, repeat int, write func(md *MarkDown, content interface{}, repeat int)) {
	md.WriteHandle(content, repeat, WriteHandlerFunc(write))
}

// var defaultMarkDown MarkDown

// var DefaultMarkDown = &defaultMarkDown

// func WriteHandleFunc(content interface{}, repeat int, write func(md *MarkDown, content interface{}, repeat int)) {
// 	DefaultMarkDown.WriteHandleFunc(content, repeat, write)
// }

// func WriteHandle(content interface{}, repeat int, write WriteHandler) {
// 	defaultMarkDown.WriteHandle(content, repeat, write)
// }
