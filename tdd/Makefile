.PHONY: run all tidy test bench help

all: tidy

## run: Run the main.go file
run:
	@echo "running main.go..."
	@go run .

## test: Run all tests
test:
	@echo "initiating tests..."
	@go test ./...

## tidy: Organize dependencies
tidy:
	@echo "organizing dependencies..."
	@go mod tidy

## bench: Run benchmarks
bench:
	@echo "running tests with benchmark..."
	@go test -bench=. ./...

## help: Show available commands
help:
	@echo "Available commands:"
	@echo "  run    - Run the main.go file"
	@echo "  test   - Run all tests"
	@echo "  tidy   - Organize dependencies"
	@echo "  bench  - Run benchmarks"
	@echo "  help   - Show this help message"
