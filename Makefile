.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


.PHONY: fmt
fmt: ## Run go fmt against code.
	gofmt -s -w .
	goimports -w .
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint
lint: ## Run go lint against code.
	golangci-lint run --skip-files ".*test.go"  -v ./...

.PHONY: gencopyright
gencopyright: ## add code copyright
	@bash hack/gencopyright.sh

.PHONY: default
default: gencopyright fmt vet lint ## Run all code ci by default: gencopyright fmt vet lint.

.PHONY: genclient
genclient: ## Gen Client Code
	hack/genclient.sh
