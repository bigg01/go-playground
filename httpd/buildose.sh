#!/usr/bin/env bash
set -x
git add main.go
git commit -m "get hostname"
git push -u origin master
oc start-build go-playgroundgit