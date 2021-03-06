REPO                  := github.com/PutskouDzmitry/golang-training-Library

PHONY: help
help: ## makefile targets description
	@echo "Usage:"
	@egrep '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##/#-/' | column -t -s "#"

.PHONY: fmt
fmt: ## automatically formats Go source code
	@echo "Running 'go fmt ...'"
	@go fmt -x "$(REPO)/..."

.PHONY: image
image: fmt ## build image from Dockerfile ./docker/server/Dockerfile
	@docker build -t kvarc/final-test-docker .

.PHONY: up
up : image ## up docker compose
	@docker-compose up -d

.PHONY: integration
integration: up
	@go test --tags=integration ./cmd/server/

.PHONY: down
down : integration
	@docker-compose down

