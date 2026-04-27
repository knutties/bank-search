.PHONY: build-index build-index-from run test

build-index:
	go run ./cmd/build-index -out ./index

build-index-from:
	@if [ -z "$(CSV)" ]; then echo "usage: make build-index-from CSV=path/to/IFSC.csv"; exit 1; fi
	go run ./cmd/build-index -csv $(CSV) -out ./index

run:
	go run .

test:
	go test -tags=unit ./...
