.PHONY: clean test
.DEFAULT_GOAL := help

##@ Utility
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make help \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


all: clean test build

tidy: ## tidy the go.mod vendor file
	go mod tidy

test: ## runs tests
	go test -v -race ./...

build: ## build this project
	go build -o line-server cmd/line-server/main.go

run: ## runs this project
	go run cmd/line-server/main.go

clean: ## removes the results of a build
	@rm ./line-server
