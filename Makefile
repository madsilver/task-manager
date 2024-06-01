GREEN=\033[36m
DEFAULT=\033[0m

.PHONY: test vendor

help: ## Display help screen
	@echo "Usage:"
	@echo "\tmake [COMMAND]"
	@echo "Commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%s\t\t$(DEFAULT)%s\n", $$1, $$2}'

run: ## Run application local
	@docker-compose up -d

test: ## Run the tests of the project
	@go test -covermode=atomic -coverprofile=coverage.out  ./...

api-doc: ## Build swagger
	@go run github.com/swaggo/swag/cmd/swag init -g ./internal/infra/server/server.go

img-api: ## Build api docker image
	@docker build -t task-manager .

img-wrk: ## Build worker docker image
	@docker build -t task-manager-worker -f DockerfileWorker .

docker: ## Run docker container
	@docker run -d --rm --net=host task-manager
	@docker run -d --rm --net=host task-manager-worker

mock: ## Build mocks
	@go get github.com/golang/mock/gomock
	@go get github.com/golang/mock/mockgen@v1.6.0
	@~/go/bin/mockgen -source=internal/adapter/core/core.go -destination=internal/adapter/core/mock/core.go