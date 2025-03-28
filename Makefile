BINARY_NAME = ezex-gateway
BUILD_DIR = build
CMD_DIR = internal/cmd/main.go

.PHONY: all fmt lint vet test unit_test race_test build_linux check clean gen-graphql

gen-graphql:
	@echo "Generating graphql code..."
	@go tool gqlgen generate ./...

fmt:
	@echo "Formatting code..."
	@go tool gofumpt -l -w .

lint:
	@echo "Running lint..."
	@go tool golangci-lint  run ./... --timeout=20m0s

unit_test:
	@echo "Running unit tests..."
	@go test ./... -v

race_test:
	@echo "Running race condition tests..."
	@go test ./... -v -race

integration_test:
	@echo "Running integration tests..."
	@go test -tags=integration ./... -v

build_linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

test: unit_test race_test

check: fmt lint

clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf $(BUILD_DIR)