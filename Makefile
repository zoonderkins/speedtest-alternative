# Project name
BINARY_NAME=speedtest-alternative

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

all: test build

build: 
	$(GOBUILD) -o build/$(BINARY_NAME) -v

test: 
	$(GOCMD) test -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

# Cross compilation
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_NAME)-linux-amd64 -v

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_NAME)-darwin-amd64 -v

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o build/$(BINARY_NAME)-darwin-arm64 -v

deps:
	$(GOGET) ./...
