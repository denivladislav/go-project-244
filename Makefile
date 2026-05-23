build:
	go build -o bin/gendiff ./cmd/gendiff

run:
	bin/gendiff $(ARGS)

lint:
	golangci-lint run -c .golangci.yml $(ARGS)

fmt:
	golangci-lint fmt -c .golangci.yml

lint-fix:
	make fmt
	make lint ARGS=--fix

test:
	go test ./...

test-coverage:
	go test ./... -coverpkg=./... -coverprofile=coverage.out -covermode=atomic

.PHONY: build run lint-install lint fmt lint-fix test test-coverage
