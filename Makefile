build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run:
	bin/hexlet-path-size

GOLANGCI_LINT := golangci-lint run -c .golangci.yml

lint: 
	$(GOLANGCI_LINT)

lintfix:
	$(GOLANGCI_LINT) --fix
