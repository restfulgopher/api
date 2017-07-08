.DEFAULT_GOAL := help

PROJECT_NAME := BETALOTEST API

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: publish
publish: ## publish images on docker hub
	@docker-compose build api test stub_api spec nginx
	@docker-compose push api test stub_api spec nginx

.PHONY: api/test
api/test: ## run api unit tests
	@docker-compose up --build test

.PHONY: api/run
api/run: ## start api, stub_api and nginx as reverse proxy
	@docker-compose build api spec stub_api
	@docker-compose up -d --build nginx

.PHONY: api/stop
api/stop: ## stop and remove services containers
	@docker-compose rm -fsv nginx api spec stub_api

.PHONY: editor/run
editor/run: ## start swagger editor container
	@docker-compose up -d editor

.PHONY: editor/stop
editor/stop: ## stop and remove swagger editor container
	@docker-compose rm -fsv editor

.PHONY: terraform/apply
terraform/apply: publish ## create remote vm with terraform and deploy services
	@terraform validate terraform/
	@terraform plan terraform/
	@terraform apply terraform/

.PHONY: terraform/destroy
terraform/destroy: ## destroy remote vm with terraform
	@terraform destroy -force terraform/
