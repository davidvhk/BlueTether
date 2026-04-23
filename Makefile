BINARY_NAME=bluetether
MAIN_FILE=bluetether.go

# Disable built-in rules and variables
MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

.PHONY: all build run clean fmt vet test help

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -v ./...

help:
	@echo "Available targets:"
	@echo "  all     : Build the binary (default)"
	@echo "  build   : Build the binary"
	@echo "  run     : Build and run the binary"
	@echo "  clean   : Remove the binary"
	@echo "  fmt     : Run go fmt on the source code"
	@echo "  vet     : Run go vet on the source code"
	@echo "  test    : Run tests"
	@echo "  help    : Show this help message"
