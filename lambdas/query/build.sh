#!/bin/bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

rm -rf bin/
mkdir bin/

cp main bin/
cp conjur-dev.pem bin/
