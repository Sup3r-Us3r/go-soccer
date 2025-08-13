# Binary name
BINARY_NAME=go-soccer

# Variables for version and commit info
VERSION=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse --short HEAD)

# Main entry point
MAIN_FILE=cmd/soccer/api/main.go
BINARY_FILE=bin/$(BINARY_NAME)

# Commands
run:
	@echo "🚀 Running application..."
	@go run $(MAIN_FILE)

start:
	@echo "🚀 Running application..."
	@./$(BINARY_FILE)

build:
	@echo "📦 Building application..."
	@go build -o bin/$(BINARY_NAME) -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT)" $(MAIN_FILE)

clean:
	@echo "🧹 Cleaning binaries..."
	@rm -rf bin

test:
	@echo "🧪 Running tests..."
	@go test ./... -v

lint:
	@echo "🔍 Running static analysis..."
	@go vet ./...

help:
	@echo "Available commands:"
	@echo "  make run     - Run application in development mode"
	@echo "  make start   - Run application in production mode"
	@echo "  make build   - Build application"
	@echo "  make clean   - Remove binaries"
	@echo "  make test    - Run tests"
	@echo "  make lint    - Run static analysis"
