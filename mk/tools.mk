# --------------------------------------------------
# Tools tooling
# --------------------------------------------------

COBRA_CLI_VERSION ?= v1.3.0
GOLANGCI_LINT_VERSION ?= v1.55.2

# Ensure curl is available
ifeq (, $(shell which curl 2> /dev/null))
$(error "'curl' is not installed or available in PATH")
endif

.PHONY: install-tools
install-tools:
	@GOBIN="$(APP_DIR)/bin" go install github.com/spf13/cobra-cli@$(COBRA_CLI_VERSION)
	@curl -sSfL \
		"https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" \
		| sh -s -- -b "$(APP_DIR)/bin" "$(GOLANGCI_LINT_VERSION)"

.PHONY: lint
lint:
	@$(APP_DIR)/bin/golangci-lint run ./...
