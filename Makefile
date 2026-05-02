.PHONY: build-index build-index-from run test clean

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
