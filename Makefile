# ==============================================================================
# Variables 

NEW_RELIC_LICENSE_KEY=*************************

# ==============================================================================
# Config-01

config-01:
	LOG_LEVEL=info \
	go run config-01/app/example/main.go
.PHONY: config-01

# ==============================================================================
# Config-02

config-02:
	NEW_RELIC_LICENSE_KEY=$(NEW_RELIC_LICENSE_KEY) \
	NEW_RELIC_APP_NAME=colombostips \
	LOG_LEVEL=info \
	go run config-02/app/example/main.go
.PHONY: config-02

# ==============================================================================
# Config-03

config-03:
	docker build \
	-f config-03/zarf/docker-compose.yml \
	-t config-03 \
	.


	# docker compose logs
.PHONY: config-03