.PHONY: run test fmt lint all

# Run a specific day: make run DAY=01
run:
	@go run ./cmd/day$(DAY) input/day$(DAY).txt

# Run all tests
test:
	@go test ./internal/...

# Run tests in short mode (skips slow tests)
test-short:
	@go test -short ./internal/...

# Format all code
fmt:
	@gofmt -w .

# Run linter
lint:
	@golangci-lint run ./...

# Run all checks
all: fmt lint test-short
