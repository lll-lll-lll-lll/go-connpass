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
q := map[string]string{"nickname": "your connpass user name"}
res, _ := client.Do(context.Background(), connpass.URL(q))
defer res.Body.Close()

var cRes connpass.Response
body, _ := io.ReadAll(res.Body)
if err := json.Unmarshal(body, &cRes); err != nil {
	log.Fatal(err)
}
fmt.Println(cRes)
```
### Example
[マークダウンに書き出した例](./example/sample.md)

[コード](./example/main.go)
