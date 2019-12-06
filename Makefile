all: build test
	:

build:
	go build -o bin/dnsmonitor cmd/main.go

test:
	go test ./...

help: build
	./bin/dnsmonitor -h

fmt:
	go fmt ./...

run: build
	./bin/dnsmonitor -domain www.google.com 

lint:
	golint ./...
