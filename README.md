### 申請
個人の利用ではIPアドレスをを1つに固定し、申請を出せば無償で利用できるので[こちら](https://help.connpass.com/api/#id4)から申請をしてみましょう。
### Install
```sh
go get github.com/lll-lll-lll-lll/go-connpass@v1.1.0
```

###  Introduction
#### simple request
```go
connpassClient := &connpass.Client{}
req := &connpass.UserRequest{}
req.SetURL(connpass.CONNPASSAPI_USER_V1 + "?")
req.NickName = []string{"your connpass nickname"}
res, err := connpassClient.Do(context.Background(), req)
if err != nil {
	log.Fatal(err)
}
defer res.Body.Close()

var cRes connpass.UserResponse
body, err := io.ReadAll(res.Body)
if err != nil {
	log.Fatal(err)
}
if err := json.Unmarshal(body, &cRes); err != nil {
	log.Fatal(err)
}
fmt.Println(cRes)
```
