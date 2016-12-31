#!/usr/bin/env bash
#go build -o main .
go get github.com/golang/glog
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t go-scratch -f Dockerfile .