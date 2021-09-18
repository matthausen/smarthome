all: help

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-z0-9A-Z_-]+:.*?## / {printf "\033[36m%-45s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

compose-live: ## compose-live
	docker-compose kill $(docker ps -q) || true
	docker-compose up --force-recreate

compose-live-detach: ## compose-detach
	docker-compose -f docker-compose-local.yml up --force-recreate -d

compose-prune:
	docker container kill $(docker ps -q) || true
	docker system prune
	docker network prune
	docker volume prune
	docker image prune

compose-kill: ## compose-kill
	docker-compose -f docker-compose.yml kill
	docker-compose -f docker-compose.yml rm --force

logs-tails: ## logs-tails
	docker-compose -f docker-compose.yml logs --follow

last-logs:  ## logs-tails
	docker-compose -f docker-compose.yml logs

.PHONY: help compose-live