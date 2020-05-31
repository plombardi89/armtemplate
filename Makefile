.PHONY: default
default: all ;

format:
	goimports -w .
	go fmt ./...

lint:
	golangci-lint run
	go vet ./...

test: lint
	go test -v ./...

test-nolint:
	go test -v ./...

all: format lint test

# build logic is encapsulated in docker for repeatability in CI environments
docker:
	docker build .

vendor:
	go mod tidy
	go mod vendor
