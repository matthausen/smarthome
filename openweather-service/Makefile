all: help

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-z0-9A-Z_-]+:.*?## / {printf "\033[36m%-45s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build-push: ## compose-live
	docker build -t matthausen/openweather-service . --no-cache
	docker push matthausen/openweather-service


.PHONY: help build-push