#!/usr/bin/env bash
#go build -o main .
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t go-scratch -f Dockerfile .