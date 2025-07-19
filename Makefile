.PHONY: all build run fmt vet lint clean

BUILD_DIR = build
BINARY_NAME = $(BUILD_DIR)/ghkit

all: fmt vet lint build

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BINARY_NAME) main.go

run:
	go run main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

clean:
	rm -rf $(BUILD_DIR)
