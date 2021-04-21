GO ?= go
PRE_COMMIT ?= pre-commit

BIN_DIR := bin

GO_FILES := $(shell find . -iname '*.go')
TOOLS_GO := tools.go

.DEFAULT_GOAL := build

.PHONY: build
build: bin/toolmgr ## Build the toolmgr binary

.PHONY: lint
lint: $(GO_FILES) ## Run linter on all files.
	$(PRE_COMMIT) run --all-files

.PHONY: test
test: ## Run all tests
	$(GO) test ./...

.PHONY: install
install: $(CMD_PKGS) ## Install the binary to the default GOBIN
	$(GO) install .

.PHONY: clean
clean: ## Clean directory.
	rm -rf $(BIN_DIR)

.PHONY: help
help: ## Show this help.
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)

$(BIN_DIR)/toolmgr: main.go
	$(GO) build -o $@ .
