.PHONY: start
# start test databases
start:
	docker-compose -f entc/integration/docker-compose.yaml up -d --scale test=0

.PHONY: stop
# stop test databases
stop:
	docker-compose -f entc/integration/docker-compose.yaml stop

.PHONY: rm
# remove test databases
rm:
	docker-compose -f entc/integration/docker-compose.yaml rm -f

.PHONY: gen
# go generate
gen:
	go generate ./...

.PHONY: test-only
# run tests only
test-only:
	go test -tags json1 ./...

.PHONY: test
# start databases if not started then run tests
test: | start test-only

.PHONY: all
# go generate, start databases if not started, run tests, stop databases and remove them
all: | gen start test-only stop rm

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
	@echo ''

.DEFAULT_GOAL := help
