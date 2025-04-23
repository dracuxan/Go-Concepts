PHONY: run all tidy main

run:
	@echo "running main.go..."
	@go run .

test:
	@echo "initiating tests..."
	@go test -v ./...

tidy:
	@echo "organizing dependencies..."
	@go mod tidy

bench:
	@echo "running tests with benchmark..."
	@go test -v -bench=. ./...
