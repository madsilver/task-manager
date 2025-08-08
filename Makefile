include .env
export

GREEN=\033[36m
DEFAULT=\033[0m
COMMIT=$(git rev-parse --short HEAD)
COVERAGE_DIR ?= .coverage
DB_URL=mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

.PHONY: test build lint sec openapi help

help: ## Display help screen
	@echo "Usage:"
	@echo "\tmake [COMMAND]"
	@echo "Commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%s\t\t$(DEFAULT)%s\n", $$1, $$2}'

secrets: jwt-sec aes-sec ## Generate all secrets and update .env

jwt-sec: ## Generate a random JWT_SECRET and update .env
	@secret=$$(openssl rand -base64 32); \
	sed -i.bak "s|^JWT_SECRET=.*|JWT_SECRET=$$secret|" .env && rm .env.bak; \

aes-sec: ## Generate a random AES-256 key and update .env
	@secret=$$(openssl rand -base64 32); \
	sed -i.bak "s|^TASK_ENCRYPTION_KEY=.*|TASK_ENCRYPTION_KEY=$$secret|" .env && rm .env.bak; \

run: ## Run the api
	@go run cmd/api/main.go

build: ## Build program
	@go build -o api cmd/api/main.go

image: ## Build docker image
	@docker build -t tm-api:latest .

migrate: ## Run database migrations using Docker
	docker run --rm -v $(PWD)/migrations:/migrations \
		--network host \
		migrate/migrate \
		-path=/migrations \
		-database "$(DB_URL)" up

openapi: ## Build open-api docs
	@go get github.com/swaggo/swag/gen@v1.8.12
	@go get github.com/swaggo/swag/cmd/swag@v1.8.12
	@go run github.com/swaggo/swag/cmd/swag init -g internal/infrastructure/server/gin_server.go -o docs/swagger

mock: ## Build mocks
	@go install go.uber.org/mock/mockgen@latest
	@mockgen -source=internal/adapter/handler/interface.go -destination=internal/adapter/handler/mock/interface.go
	@mockgen -source=internal/usecase/user_usecase.go -destination=internal/usecase/mock/user_usecase.go
	@mockgen -source=internal/usecase/task_usecase.go -destination=internal/usecase/mock/task_usecase.go

test: ## Run the tests of the project
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	@go test -v -covermode atomic -coverprofile $(COVERAGE_DIR)/coverage.txt $(shell go list ./... | grep -vE '/(mock|infrastructure|cmd|docs)')

html-c: ## Generate HTML coverage report
	go tool cover -html=$(COVERAGE_DIR)/coverage.txt

lint: ## Run linter
	@golangci-lint run -c .golangci.yml ./...

sec: ## Run vulnerability scan
	@govulncheck -scan symbol -show verbose ./...

