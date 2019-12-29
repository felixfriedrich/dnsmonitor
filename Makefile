all: build generate test tidy update
	:

build:
	go build -o bin/dnsmonitor cmd/main.go

release: fmt lint update test
	env GOOS=linux GOARCH=amd64 go build -o bin/dnsmonitor-linux-amd64 cmd/main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/dnsmonitor-darwin-amd64 cmd/main.go

update:
	go get -u -t ./...

tidy:
	go mod tidy && go generate ./...
	# go mod tidy removes couterfeiter from dependencies
	# go generate adds it again

generate:
	go generate ./...

test:
	go test ./...

help: build
	./bin/dnsmonitor -h

fmt:
	go fmt ./...

run: build
	./bin/dnsmonitor -domain www.google.com 

lint:
	golint -set_exit_status ./...
