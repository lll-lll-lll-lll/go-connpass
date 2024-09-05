### 申請
個人の利用ではIPアドレスをを1つに固定し、申請を出せば無償で利用できるので[こちら](https://help.connpass.com/api/#id4)から申請をしてみましょう。
### Install
```sh
go get github.com/lll-lll-lll-lll/go-connpass@v1.1.0
```

###  Introduction
#### simple request
```go
client := &connpass.Client{}
req := &connpass.EventRequest{}
req.SetURL(connpass.CONNPASSAPI_EVENT_V1 + "?")
req.NickName = []string{"your connpass nickname"}
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
res, _ := client.Do(ctx, req)
defer res.Body.Close()

var cRes connpass.EventResponse
body, _ := io.ReadAll(res.Body)
if err := json.Unmarshal(body, &cRes); err != nil {
	log.Fatal(err)
}
fmt.Println(cRes)
```
