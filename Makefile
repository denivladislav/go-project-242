build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run:
	bin/hexlet-path-size $(ARGS)

lint:
	golangci-lint run -c .golangci.yml $(ARGS)

fmt:
	golangci-lint fmt -c .golangci.yml

lint-fix:
	make fmt
	make lint ARGS=--fix

test:
	go test ./...

.PHONY: build run lint fmt lint-fix test
