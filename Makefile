.PHONY: generate clean

generate:
	@echo "Generating GraphQL code..."
	@go generate ./...
	# @go run github.com/99designs/gqlgen generate

clean:
	@echo "Cleaning generated GraphQL files..."
	@rm -rf internal/adapters/handlers/graphql/generated
