.PHONY: help build test run clean

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build           Build the calculator binary"
	@echo "  test            Run tests"
	@echo "  test-coverage   Run tests with code coverage"
	@echo "  run             Run the calculator"
	@echo "  clean           Clean build artifacts"
	@echo "  all             Build and run the calculator"
	@echo "  help            Show this help message"

# Build the calculator binary
build:
	@echo "Building calculator..."
	@cd src && go build -o ../bin/calculator main.go
	@echo "Build complete: bin/calculator"

# Run tests
test:
	@echo "Running tests..."
	@cd src && go test -v

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@cd src && go test -v -cover

# Run the calculator
run:
	@echo "Running calculator..."
	@cd src && go run main.go

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf bin/
	@echo "Clean complete"

# Build and run
all: build run
