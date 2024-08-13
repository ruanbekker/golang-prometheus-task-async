.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# Determine whether to use "docker-compose" or "docker compose"
DOCKER_COMPOSE := $(shell which docker-compose 2>/dev/null)
ifeq ($(DOCKER_COMPOSE),)
	DOCKER_COMPOSE := $(shell which docker 2>/dev/null)
	PREFIX := compose
else
	PREFIX :=
endif

# DOCKER TASKS
build: ## Builds docker image
	$(DOCKER_COMPOSE) $(PREFIX) build

up: ## Runs the containers in detached mode with default config
	$(DOCKER_COMPOSE) $(PREFIX) up -d --build

clean: ## Stops and removes all containers
	$(DOCKER_COMPOSE) $(PREFIX) -f docker-compose.yml -f down

logs: ## View the logs from the containers
	$(DOCKER_COMPOSE) $(PREFIX) logs -f


