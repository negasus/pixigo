.PHONY: help gen build-wasm build

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gen: ## Go generate
	@go generate ./...

build-wasm:
	@echo "Building wasm file for example: $(NAME)"
	@GOOS=js GOARCH=wasm go build -o examples/basic/$(NAME)/main.wasm examples/basic/$(NAME)/main.go

build: ## Build examples wasm files
	@make build-wasm NAME=1-container
	@make build-wasm NAME=3-tinting
	@make build-wasm NAME=4-particle-container

