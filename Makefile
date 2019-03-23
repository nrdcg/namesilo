.PHONY: clean check test build dependencies fmt

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

default: clean check test build

clean:
	rm -rf cover.out

build: clean
	GO111MODULE=on go build -v .

dependencies:
	GO111MODULE=on go mod download

test: clean
	GO111MODULE=on go test -v -cover ./...

check:
	GO111MODULE=on golangci-lint run

fmt:
	gofmt -s -l -w $(SRCS)

## Useful to initiate structures
gen-struct:
	echo 'package namesilo' > "gen_struct.go";
	echo '' >> "gen_struct.go";
	echo 'import "encoding/xml"' >> "gen_struct.go";
	echo '' >> "gen_struct.go";
	for i in $$(ls samples/ -1); do \
		zek -c -n $${i%.xml} "./samples/$$i" >> "gen_struct.go"; \
	done
