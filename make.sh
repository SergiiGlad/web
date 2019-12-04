#!/usr/bin/env bash
#
# build for dockerize container
#

CGO_ENABLED=0 go build -a -installsuffix cgo -o ${GOPATH}/bin web/wiki.go
