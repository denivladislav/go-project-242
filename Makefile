build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run:
	bin/hexlet-path-size $(ARGS)

lint-install:
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.11.2

lint:
	golangci-lint run -c .golangci.yml $(ARGS)

fmt:
	golangci-lint fmt -c .golangci.yml

lint-fix:
	make fmt
	make lint ARGS=--fix

test:
	go test ./...

.PHONY: build run lint-install lint fmt lint-fix test
