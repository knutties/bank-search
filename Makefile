.PHONY: build-index build-index-from run test clean \
        smithy-build smithy-publish smithy-updates smithy-clean

build-index:
	go run ./cmd/build-index -out ./index

build-index-from:
	@if [ -z "$(CSV)" ]; then echo "usage: make build-index-from CSV=path/to/IFSC.csv"; exit 1; fi
	go run ./cmd/build-index -csv $(CSV) -out ./index

run:
	@if [ ! -d ./index ]; then \
	  echo "no index at ./index — run 'make build-index' (or 'make build-index-from CSV=...') first"; \
	  exit 1; \
	fi
	go run .

test:
	go test -tags=unit ./...

clean:
	rm -rf ./index

# Builds Smithy projections into smithy/build/.
smithy-build:
	cd smithy && smithy build

# Copies the generated OpenAPI spec and Smithy AST into api/ so consumers
# don't need a smithy-cli to read them.
smithy-publish: smithy-build
	cp smithy/build/smithy/source/openapi/BankSearch.openapi.json api/openapi.json
	cp smithy/build/smithy/source/model/model.json api/model.json

smithy-clean:
	rm -rf smithy/build

# Used by CI to assert the committed api/ artifacts match the IDL.
smithy-updates: smithy-clean smithy-publish
