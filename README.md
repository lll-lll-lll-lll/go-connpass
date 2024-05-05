### Install
```sh
go get github.com/lll-lll-lll-lll/go-connpass
```

### Instruction
#### simple request
```go
client := &connpass.Client{}
q := map[string]string{"nickname": "your connpass user name"}
res, _ := client.Do(context.Background(), connpass.URLV1(q))
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