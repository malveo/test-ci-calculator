.PHONY: help build test run clean docker-build docker-run

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
	@echo "  docker-build    Build Docker image"
	@echo "  docker-run      Run calculator in Docker (example: make docker-run ARGS='add 10 5')"
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

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -f docker/Dockerfile -t calculator:latest .
	@echo "Docker image built: calculator:latest"

# Run calculator in Docker
docker-run:
	@echo "Running calculator in Docker..."
	@docker run --rm calculator:latest $(ARGS)
