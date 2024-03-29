GREEN=\033[36m
DEFAULT=\033[0m

.PHONY: test vendor

help: ## Display help screen
	@echo "Usage:"
	@echo "\tmake [COMMAND]"
	@echo "Commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%s\t\t$(DEFAULT)%s\n", $$1, $$2}'

run: ## Run application
	@docker-compose up -d
	@go run cmd/api/main.go

tidy: ## Downloads go dependencies
	@go mod tidy

vendor: ## Copy of all packages needed
	@go mod vendor

test: ## Run the tests of the project
	@go test -covermode=atomic -coverprofile=coverage.out  ./...

test-v: ## Run the tests of the project (verbose)
	@go test -v -cover -p=1 -covermode=count -coverprofile=coverage.out  ./...

mock: ## Build mocks
	@go get github.com/golang/mock/gomock
	@go get github.com/golang/mock/mockgen@v1.6.0
	@~/go/bin/mockgen -source=internal/adapter/controller/task.go -destination=internal/adapter/controller/mock/task.go
	@~/go/bin/mockgen -source=internal/adapter/repository/mysql/task.go -destination=internal/adapter/repository/mysql/mock/task.go