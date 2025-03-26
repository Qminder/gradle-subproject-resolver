#!/usr/bin/env bash
set -e

GOOS=darwin GOARCH=amd64 go build -o gradle-subproject-resolver-macos-amd64 *.go
GOOS=darwin GOARCH=arm64 go build -o gradle-subproject-resolver-macos-arm64 *.go
lipo -create -output gradle-subproject-resolver-macos gradle-subproject-resolver-macos-amd64 gradle-subproject-resolver-macos-arm64

rm -f gradle-subproject-resolver-macos-amd64 gradle-subproject-resolver-macos-arm64

GOOS=linux GOARCH=amd64 go build -o gradle-subproject-resolver-linux-amd64 *.go
GOOS=linux GOARCH=arm64 go build -o gradle-subproject-resolver-linux-arm64 *.go
