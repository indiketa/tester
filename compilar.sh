#!/bin/bash

GOPATH=/home/ecatala/Codi/gopath
docker run --rm -v "$GOPATH":/go -v "$PWD":/out golang:1.9-alpine3.7 go build -v -i -o /out/servidor src/docker-test/main.go
docker build . -t indiketa/tester:1.0
