# Variables
APP_NAME = mkv-cli

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test

# Determine the installation directory
GOBIN_DIR = $(shell go env GOBIN)
ifeq ($(GOBIN_DIR),)
  GOBIN_DIR = $(shell go env GOPATH)/bin
endif

# Build flags to strip debug information and reduce binary size
LDFLAGS = -w -s

.PHONY: all
all: build

# Build the project
.PHONY: build
build: 
	$(GOBUILD) -ldflags="$(LDFLAGS)" -o $(APP_NAME) main.go

# Build and install the project
.PHONY: install
install:
	$(GOINSTALL) -ldflags="$(LDFLAGS)" ./...

# Run the tests
.PHONY: test
test: 
	$(GOTEST) -v ./...

# Uninstall the project
.PHONY: uninstall
uninstall:
	rm -f $(GOBIN_DIR)/$(APP_NAME)

# Clean the build directory
.PHONY: clean
clean: 
	$(GOCLEAN)
	rm -f $(APP_NAME)

# Format the code
.PHONY: fmt
fmt:
	gofmt -s -w .

# Display help
.PHONY: help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo
	@echo "Usage:"
	@echo "  make [target]"
	@echo
	@echo "Targets:"
	@echo "  all           Default target: build"
	@echo "  build         Build the project"
	@echo "  install       Install the binary"
	@echo "  uninstall     Uninstall the binary"
	@echo "  clean         Clean the build directory"
	@echo "  fmt           Format the code"
	@echo "  test          Run tests"
	@echo "  help          Display this help message"