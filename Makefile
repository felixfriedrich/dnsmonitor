all: clean build generate test tidy update
	:

build:
	go build -o bin/dnsmonitor cmd/dnsmonitor/main.go

run-integration-test:
	go run cmd/integration_test/main.go

release: all
	env GOOS=linux GOARCH=amd64 go build -o bin/dnsmonitor-linux-amd64 cmd/dnsmonitor/main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/dnsmonitor-darwin-amd64 cmd/dnsmonitor/main.go

update:
	go get -u -t ./...

tidy:
	go mod tidy && go generate ./...
	# go mod tidy removes counterfeiter from dependencies
	# go generate adds it again

clean:
	go clean -testcache
	rm -f test.out
	rm -f ./bin/*


generate:
	find . -name "*fakes" -exec rm -rf -- {} + && go generate ./...

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
