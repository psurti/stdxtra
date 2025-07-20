# Makefile for stdxtra Go project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

# Project name
PROJECT_NAME=stdxtra

# Main package path
MAIN_PACKAGE=stdxtra

# Build directory
BUILD_DIR=./build

# Binary name (automatically adds .exe on Windows)
BINARY_NAME=$(PROJECT_NAME)

# Default target
.PHONY: all
all: clean fmt test build

# Build the project
.PHONY: build
build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) .

# Test the project
.PHONY: test
test:
	$(GOTEST) ./...

# Clean build artifacts
.PHONY: clean
clean:
	$(GOCLEAN) ./..
# Format the code
.PHONY: fmt
fmt:
	$(GOFMT) ./...

# Check if code is formatted without changing it
.PHONY: fmt-check
fmt-check:
	@$(GOCMD) run -e "package main; import ( \"fmt\"; \"os\"; \"os/exec\" ); func main() { cmd := exec.Command(\"$(GOFMT)\", \"-l\", \"./...\"); output, _ := cmd.Output(); if len(output) > 0 { fmt.Println(\"The following files are not formatted:\"); fmt.Println(string(output)); os.Exit(1) } else { fmt.Println(\"All files are properly formatted.\") } }"

# Run tests with verbose output
.PHONY: test-verbose
test-verbose:
	$(GOTEST) -v ./...

# Run linting tools
.PHONY: lint
lint:
	@$(GOCMD) run -e "package main; import ( \"fmt\"; \"os\"; \"os/exec\" ); func main() { _, err := exec.LookPath(\"golangci-lint\"); if err == nil { cmd := exec.Command(\"golangci-lint\", \"run\", \"./...\"); cmd.Stdout = os.Stdout; cmd.Stderr = os.Stderr; cmd.Run() } else { fmt.Println(\"golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest\") } }"

# Update dependencies
.PHONY: deps
deps:
	$(GOMOD) tidy

# Run the application
.PHONY: run
run:
	@$(GOCMD) run main.go

# Run benchmarks
.PHONY: bench
bench:
	$(GOTEST) -bench=. ./...

# Generate code coverage report
.PHONY: coverage
coverage:
	$(GOTEST) -coverprofile=$(BUILD_DIR)/coverage.out ./...
	$(GOCMD) tool cover -html=$(BUILD_DIR)/coverage.out -o $(BUILD_DIR)/coverage.html
	@echo "Coverage report generated at $(BUILD_DIR)/coverage.html"

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all          - Run clean, fmt, test, and build"
	@echo "  build        - Build the project"
	@echo "  test         - Run tests"
	@echo "  test-verbose - Run tests with verbose output"
	@echo "  bench        - Run benchmarks"
	@echo "  coverage     - Generate code coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  fmt          - Format code"
	@echo "  fmt-check    - Check if code is properly formatted"
	@echo "  lint         - Run linters"
	@echo "  deps         - Update dependencies"
	@echo "  run          - Build and run the application"
	@echo "  help         - Show this help message"
