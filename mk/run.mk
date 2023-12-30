# --------------------------------------------------
# Run tooling
# --------------------------------------------------

.PHONY: license
license:
	@go run "$(APP_DIR)/main.go" license

.PHONY: migrate-down
migrate-down: build
	@"$(APP_DIR)/bin/tpl" migrate down

.PHONY: migrate-redo
migrate-redo: build
	@"$(APP_DIR)/bin/tpl" migrate redo

.PHONY: migrate-reset
migrate-reset: build
	@"$(APP_DIR)/bin/tpl" migrate reset

.PHONY: migrate-status
migrate-status: build
	@"$(APP_DIR)/bin/tpl" migrate status

.PHONY: migrate-up
migrate-up: build
	@"$(APP_DIR)/bin/tpl" migrate up

.PHONY: version
version: build
	@"$(APP_DIR)/bin/tpl" version
