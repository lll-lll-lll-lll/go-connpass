package markdown

type MDElement struct {
	handler WriteHandler
}

func (mde *MDElement) WriteHandler(write WriteHandler) *MDElement {
	return mde
}

func (mde *MDElement) WriteHandlerFunc(write func(content interface{}, repeat int)) *MDElement {
	return mde.WriteHandler(WriteHandlerFunc(write))
}
