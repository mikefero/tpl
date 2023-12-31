# --------------------------------------------------
# Tools tooling
# --------------------------------------------------

COBRA_CLI_VERSION ?= v1.3.0
GINKGO_VERSION ?= v2.13.2
GOLANGCI_LINT_VERSION ?= v1.55.2

# Ensure curl is available
ifeq (, $(shell which curl 2> /dev/null))
$(error "'curl' is not installed or available in PATH")
endif

# Ensure docker is available
ifeq (, $(shell which docker 2> /dev/null))
$(error "'docker' is not installed or available in PATH")
endif

.PHONY: install-tools
install-tools:
	@GOBIN="$(APP_DIR)/bin" go install github.com/spf13/cobra-cli@$(COBRA_CLI_VERSION)
	@GOBIN="$(APP_DIR)/bin" go install github.com/onsi/ginkgo/v2/ginkgo@$(GINKGO_VERSION)
	@curl -sSfL \
		"https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" \
		| sh -s -- -b "$(APP_DIR)/bin" "$(GOLANGCI_LINT_VERSION)"

.PHONY: lint-oas
lint-oas:
	@docker run --rm -it \
		-v $(APP_DIR)/openapi.yaml:/spec/openapi.yaml \
		-v $(APP_DIR)/redocly.yaml:/spec/redocly.yaml \
		redocly/cli lint \
			--config /spec/redocly.yaml \
			/spec/openapi.yaml

.PHONY: lint
lint: list-oas
	@$(APP_DIR)/bin/golangci-lint run ./...
