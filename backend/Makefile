# Makefile for Customer Family API

.PHONY: docs run build clean

# Generate Swagger documentation
docs:
	swag init -g cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Build the application
build:
	go build -o bin/app cmd/main.go

# Clean generated files
clean:
	rm -rf bin/
	rm -rf docs/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run with live reload (requires fresh)
dev:
	fresh

# Generate docs and run
docs-run: docs run

# Help
help:
	@echo "Available commands:"
	@echo "  docs      - Generate Swagger documentation"
	@echo "  run       - Run the application"
	@echo "  build     - Build the application"
	@echo "  clean     - Clean generated files"
	@echo "  deps      - Install/update dependencies"
	@echo "  dev       - Run with live reload using Fresh"
	@echo "  docs-run  - Generate docs and run"
	@echo "  help      - Show this help message"
