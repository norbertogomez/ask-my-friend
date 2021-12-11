SHELL := /bin/bash

# Makefile variables
default_network := amf_default
users-app-container := users-app

help: ## Displays this help ❓
	@echo 'Usage: make [target] ...'
	@echo
	@echo 'targets:'
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"

logs-amf-users:
	@docker-compose -f users/docker-compose.yaml logs

shell-amf-users: ## Starts a shell for users container
	@docker-compose -f users/docker-compose.yaml exec amf-users sh

start: start-amf-users ## Starts the application 🎬.

start-amf-users: ## Starts users app with docker compose🏃
	@echo -e "🏃 - Running Users App - 🏃\n"
	@docker-compose -f users/docker-compose.yaml up  --build -d
	@echo -e "\n✅ - Done - ✅\n"

status: ## Displays the status of all the services 💤.
	@echo -e "🐳 - Docker compose containers - 🐳\n"
	@docker-compose -f users/docker-compose.yaml ps

stop: stop-amf-users ## Stops the application and all it's components. 🛑

stop-amf-users: ## Stops users app 🛑
	@echo -e "🛑 - Stopping the users app - 🛑\n"
	@docker-compose -f users/docker-compose.yaml down
	@echo -e "\n✅ - Done - ✅\n"

.DEFAULT_GOAL := help