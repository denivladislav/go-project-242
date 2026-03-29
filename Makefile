.PHONY: build
build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

.PHONY: run
run:
	bin/hexlet-path-size $(ARGS)

GOLANGCI_LINT = golangci-lint
GOLANGCI_CONFIG_PATH = .golangci.yml

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run -c $(GOLANGCI_CONFIG_PATH) $(ARGS)

.PHONY: fmt
fmt:
	$(GOLANGCI_LINT) fmt -c $(GOLANGCI_CONFIG_PATH) 

.PHONY: lintfix
lint-fix:
	make fmt
	make lint ARGS=--fix

.PHONY: test
test:
	go test