#!/usr/bin/env bash

mkdir -p dist

env GOOS=linux GOARCH=amd64 go build -mod=vendor -o dist/go-sample-api-linux