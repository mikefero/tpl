# --------------------------------------------------
# Test tooling
# --------------------------------------------------

.PHONY: test
test:
	@go test -race ./...