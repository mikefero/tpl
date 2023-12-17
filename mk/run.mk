# --------------------------------------------------
# Run tooling
# --------------------------------------------------

.PHONY: license
license:
	@go run "$(APP_DIR)/main.go" license

.PHONY: version
version: build
	@"$(APP_DIR)/bin/tpl" version
