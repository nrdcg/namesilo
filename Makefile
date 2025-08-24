.PHONY: clean format check test build

export GO111MODULE=on

default: clean format check test build

clean:
	rm -rf cover.out

build: clean
	 go build -v .

test: clean
	CGO_ENABLED=1 go test -v -cover -race -json -count=1 ./... | tparse -all

format:
	gofumpt -l -w -e -extra .

check:
	golangci-lint run --timeout=5m --max-same-issues=0

## Useful to initiate structures
gen-struct:
	echo 'package namesilo' > "gen_struct.go";
	echo '' >> "gen_struct.go";
	echo 'import "encoding/xml"' >> "gen_struct.go";
	echo '' >> "gen_struct.go";
	for i in $$(ls samples/ -1); do \
		zek -c -n $${i%.xml} "./samples/$$i" >> "gen_struct.go"; \
	done
