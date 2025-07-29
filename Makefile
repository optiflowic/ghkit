.PHONY: all build run fmt vet lint clean

BUILD_DIR = build
BINARY_NAME = $(BUILD_DIR)/ghkit
VERSION ?= $(shell git describe --tags --always --dirty)

install:
	go install go.uber.org/mock/mockgen@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/segmentio/golines@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

all: generate fmt vet lint test build

build:
	mkdir -p $(BUILD_DIR)
	go build -ldflags "-X github.com/optiflowic/ghkit/cmd.version=$(VERSION)" -o $(BINARY_NAME) ./main.go

run:
	go run ./main.go

fmt:
	go fmt ./...
	golines --base-formatter=gofmt -w .

vet:
	go vet ./...

lint:
	golangci-lint run ./...
	gosec -quiet ./...

test:
	go test ./...

generate:
	go generate ./...

clean:
	rm -rf $(BUILD_DIR)
