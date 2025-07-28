.PHONY: all build run fmt vet lint clean

BUILD_DIR = build
BINARY_NAME = $(BUILD_DIR)/ghkit
VERSION ?= $(shell git describe --tags --always --dirty)

all: fmt vet lint test build

build:
	mkdir -p $(BUILD_DIR)
	go build -ldflags "-X github.com/optiflowic/ghkit/cmd.version=$(VERSION)" -o $(BINARY_NAME) ./main.go

run:
	go run main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

test:
	go test ./...

clean:
	rm -rf $(BUILD_DIR)
