# --------------------------------------------------
# Test tooling
# --------------------------------------------------

.PHONY: test
test:
	@go test -race ./...

.PHONY: test-ginkgo
test-ginkgo:
	@"$(APP_DIR)/bin/ginkgo" --race --cover --randomize-all ./...
