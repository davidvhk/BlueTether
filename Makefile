BINARY_NAME=bluetether
MAIN_FILE=bluetether.go

# Disable built-in rules and variables
MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

.PHONY: all build run clean fmt vet test deb rpm package help

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
	rm -f *.deb *.rpm

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -v ./...

deb: build
	nfpm pkg --target bluetether.deb

rpm: build
	nfpm pkg --target bluetether.rpm

package: deb rpm

help:
	@echo "Available targets:"
	@echo "  all     : Build the binary (default)"
	@echo "  build   : Build the binary"
	@echo "  run     : Build and run the binary"
	@echo "  clean   : Remove the binary and packages"
	@echo "  fmt     : Run go fmt on the source code"
	@echo "  vet     : Run go vet on the source code"
	@echo "  test    : Run tests"
	@echo "  deb     : Build DEB package (requires nfpm)"
	@echo "  rpm     : Build RPM package (requires nfpm)"
	@echo "  package : Build both DEB and RPM packages"
	@echo "  help    : Show this help message"
