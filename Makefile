.PHONY: generate clean

generate:
	@echo "Generating GraphQL code..."
	@go generate ./...
	# @go run github.com/99designs/gqlgen generate

clean:
	@echo "Cleaning generated GraphQL files..."
	@rm -rf internal/adapters/handlers/graphql/generated

test:
	@echo "Running tests..."
	@go test ./...

start:
	@echo "Starting server..."
	@go run cmd/main.go
