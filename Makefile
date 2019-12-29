all: clean build generate test tidy update
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
	# go mod tidy removes counterfeiter from dependencies
	# go generate adds it again

clean:
	go clean -testcache
	find . -name "*fakes" -exec rm -rf -- {} + # delete all generated mocks
	rm -f test.out

generate:
	go generate ./...

test:
	go test -cover ./...

test-report:
	go test ./... -coverprofile=test.out
	go tool cover -html=test.out

help: build
	./bin/dnsmonitor -h

fmt:
	go fmt ./...

run: build
	./bin/dnsmonitor -domain www.google.com 

lint:
	golint -set_exit_status ./...
