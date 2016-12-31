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


PROTO

https://github.com/golang/protobuf

http://www.grpc.io/




openshift

https://github.com/openshift-s2i/s2i-go.git \
    --context-dir=1.4/test/test-app openshift/go-14-centos7
    
    
    https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/