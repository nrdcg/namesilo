.PHONY: clean check test build

export GO111MODULE=on

default: clean check test build

clean:
	rm -rf cover.out

build: clean
	 go build -v .

test: clean
	go test -v -cover ./...

check:
	golangci-lint run

## Useful to initiate structures
# TODO: must be updated to support the subfolders.
#gen-struct:
#	# go install github.com/miku/zek/cmd/zek@latest
#	echo 'package namesilo' > "gen_struct.go";
#	echo '' >> "gen_struct.go";
#	echo 'import "encoding/xml"' >> "gen_struct.go";
#	echo '' >> "gen_struct.go";
#	for i in $$(ls samples/ -1); do \
#		zek -c -n $${i%.xml} "./samples/$$i" >> "gen_struct.go"; \
#	done

generate:
	go run ./internal/
