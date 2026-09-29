.PHONY: help test lint

help:
	@echo "Runphase"
	@echo "  make test   run Go unit tests"
	@echo "  make lint   run golangci-lint"

test:
	go test ./...

lint:
	golangci-lint run ./...
