# Weight Controller Makefile

BINARY_NAME := weight-controller
BUILD_DIR := bin

IMG_REGISTRY ?= ghcr.io/yehuditkerido
IMG_NAME ?= weight-controller
IMG_TAG ?= latest
IMG ?= $(IMG_REGISTRY)/$(IMG_NAME):$(IMG_TAG)

GOLANGCI_LINT_VERSION ?= v2.6.2

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

##@ Development

.PHONY: build
build: tidy lint test binary ## Full build: tidy, lint, test, binary

.PHONY: binary
binary: ## Build binary
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd

.PHONY: run
run: binary ## Run locally
	./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: test
test: ## Run tests
	go test -race -coverprofile=coverage.out ./...

.PHONY: tidy
tidy: ## Run go mod tidy
	go mod tidy

.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run

##@ Container

.PHONY: docker-build
docker-build: ## Build container image
	docker build -t $(IMG) .

.PHONY: docker-push
docker-push: ## Push container image
	docker push $(IMG)

##@ Deployment

.PHONY: install
install: ## Deploy to cluster
	kubectl apply -k deploy/

.PHONY: uninstall
uninstall: ## Remove from cluster
	kubectl delete -k deploy/ --ignore-not-found

##@ Clean

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf $(BUILD_DIR) coverage.out
