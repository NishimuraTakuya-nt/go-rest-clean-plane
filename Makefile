.PHONY: lint lint-fix build up clean help

NAME := go-rest-clean-plane
DC := docker compose
LDFLAGS := -ldflags="-s -w -extldflags \"-static\""

## Local ##
lint: ## Run linter
	golangci-lint run

lint-fix: ## Run linter with fix
	golangci-lint run --fix

test: ## Run tests
	go test -v ./...

## Container ##
build: ## Build image
	$(DC) build $(NAME)

up: ## Run container
	$(DC) up -d

clean: ## Clean up
	docker system prune -f

help: ## display this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
