https://golang.org/doc/code.html

gofmt -w stringutil/reverse.go
  543  gofmt -w hello/hello.go
  544  go env


go install github.com/bigg01/go-playground/hello

$GOPATH/bin/hello
gofmt -w stringutil/reverse.go
gofmt -w hello/hello.go
go env
go build stringutil/
go install hello/
$GOPATH/bin/hello

```
go test
PASS
ok  	github.com/bigg01/go-playground/stringutil	0.005s
```