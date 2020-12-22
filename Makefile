.PHONY: test build

all: clean fmt lint code-check imports build generate test tidy update
	echo ":-)"

build:
	go build -o bin/dnsmonitor cmd/dnsmonitor/main.go

release: all
	env GOOS=linux GOARCH=amd64 go build -o bin/dnsmonitor-linux-amd64 cmd/dnsmonitor/main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/dnsmonitor-darwin-amd64 cmd/dnsmonitor/main.go
	cd bin && sha256sum dnsmonitor-* > ../SHA256SUM.txt

update:
	go get -u -t ./...

tidy:
	go mod tidy
	# go mod tidy removed the following packages
	# They need to be reinstalled
	go get honnef.co/go/tools/cmd/staticcheck
	go get github.com/maxbrunsfeld/counterfeiter/v6

clean:
	go clean -testcache
	rm -f test.out
	rm -f ./bin/*
	rm -rf ./build/bin
	#docker rmi -f dnsmonitor

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

run-config-file: build
	./bin/dnsmonitor -configfile configs/default.yml

lint:
	golint -set_exit_status ./...

code-check:
	ineffassign .

imports:
	find . -name "*.go" -exec goimports -w {} \;

unused:
	staticcheck -unused.whole-program ./...

docker:
	./build/docker-build.sh

docker-push:
	./build/docker-push.sh
