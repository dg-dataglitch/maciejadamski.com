.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: templ
templ: ## Generate templ templates
	@templ generate

.PHONY: fmt
fmt: ## Format templ templates
	@templ fmt .

.PHONY: build
build: templ ## Generate static site
	@echo "> Cleaning dist..."
	@rm -rf dist
	@echo "> Building static site..."
	@go run ./cmd/build
	@echo "> Build complete!"

.PHONY: dev
dev: ## Run development server with hot reload
	@echo "> Starting dev server with Air..."
	@echo "> Local:   http://localhost:3000"
	@echo "> Network: http://$$(hostname -I | awk '{print $$1}'):3000"
	@air

.PHONY: clean
clean: ## Clean build artifacts
	@rm -rf dist tmp

.PHONY: docker
docker: ## Build Docker image
	@docker build -t go-ssg-template .

.PHONY: docker-run
docker-run: docker ## Run Docker container
	@docker run -p 8080:80 go-ssg-template
