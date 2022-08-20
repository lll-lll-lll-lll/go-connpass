# 使い方
1. README.mdに書き込みたい内容を用意する
2. 記法と何回その記法を繰り返すかを設定("#", 2 => "##")
3. それをWriteHandleFuncに渡す

例
```go
package main

import (
	"log"
	"os"

	"github.com/conread/markdown"
)

func WriteHorizon(m *markdown.MarkDown, content interface{}, repeat int) {
	markh := "-"
	m.AddToPage(markh, content, repeat)
}

func WriteTitle(m *markdown.MarkDown, content interface{}, repeat int) {
	markt := "#"
	m.AddToPage(markt, content, repeat)
}

func main() {
	file, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	m := markdown.NewMarkDown()
	m.WriteHandleFunc("Test Write Title", 2, WriteTitle)
	m.WriteHandleFunc("Test Write Horizon", 3, WriteHorizon)
	s := m.CompleteMDFile(2)
	file.Write([]byte(s))
}
```

## 結果<br>
README.md
```md
## Test Write Title

--- Test Write Horizon
```