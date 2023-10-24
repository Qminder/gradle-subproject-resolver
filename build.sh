#!/usr/bin/env bash
set -e

GOOS=darwin GOARCH=amd64 go build -o gradle-subproject-resolver-amd64 *.go
GOOS=darwin GOARCH=arm64 go build -o gradle-subproject-resolver-arm64 *.go
lipo -create -output gradle-subproject-resolver gradle-subproject-resolver-amd64 gradle-subproject-resolver-arm64

rm -f gradle-subproject-resolver-amd64 gradle-subproject-resolver-arm64
