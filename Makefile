.PHONY: clean check test build dependencies fmt

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

default: clean check test build

clean:
	rm -rf cover.out

build: clean
	go build -v .

dependencies:
	dep ensure -v

test: clean
	go test -v -cover ./...

check:
	golangci-lint run

fmt:
	gofmt -s -l -w $(SRCS)
