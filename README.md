### Install
```sh
go get github.com/lll-lll-lll-lll/go-connpass
```

### Instruction
#### simple request
```go
client := &connpass.Client{}
q := map[string]string{"nickname": "your connpass user name"}
res, _ := client.Do(connpass.Query(q), connpass.URLV1())
defer res.Body.Close()

var cRes connpass.Response
body, _ := io.ReadAll(res.Body)
if err := json.Unmarshal(body, &cRes); err != nil {
	log.Fatal(err)
}
fmt.Println(cRes)
```

