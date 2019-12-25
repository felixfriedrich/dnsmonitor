all: build test
	:

build:
	go build -o bin/dnsmonitor cmd/main.go

release: fmt lint update test
	env GOOS=linux GOARCH=amd64 go build -o bin/dnsmonitor-linux-amd64 cmd/main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/dnsmonitor-darwin-amd64 cmd/main.go

update:
	go get -u -t ./...

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
