# --------------------------------------------------
# Docker tooling
# --------------------------------------------------

# Ensure docker is available
ifeq (, $(shell which docker 2> /dev/null))
$(error "'docker' is not installed or available in PATH")
endif

# Ensure docker compose is available
ifeq (, $(shell which docker compose 2> /dev/null))
	$(error "'docker compose' is not installed or available in PATH")
endif

.PHONY: postgres-down
postgres-down:
	@docker compose -f $(APP_DIR)/docker/docker-compose-postgres.yaml down

.PHONY: postgres-up
postgres-up:
	@docker compose -f $(APP_DIR)/docker/docker-compose-postgres.yaml up -d
