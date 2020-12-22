#!/usr/bin/env bash
set -e
set -x

source $(dirname ${BASH_SOURCE})/.common.sh

env CGO_ENABLED=0 go build -a -installsuffix cgo -o ./build/dnsmonitor ./cmd/dnsmonitor/main.go

docker build \
  -t dnsmonitor${GOOS}-${GOARCH}:latest \
  -t docker.pkg.github.com/felixfriedrich/dnsmonitor/dnsmonitor-${GOOS}-${GOARCH}:latest \
  -t ghcr.io/felixfriedrich/dnsmonitor-${GOOS}-${GOARCH} \
  ./build
