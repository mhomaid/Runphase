.PHONY: help test lint fmt vet check generate dev dev-down api worker web migrate

help:
	@echo "Runphase"
	@echo "  make help       list targets"
	@echo "  make test       run go test ./..."
	@echo "  make lint       run golangci-lint"
	@echo "  make fmt        rewrite Go sources with gofmt -w"
	@echo "  make vet        run go vet ./..."
	@echo "  make check      run fmt, vet, test, and lint"
	@echo "  make generate   run go generate ./..."
	@echo "  make dev        placeholder; services start in a later task"
	@echo "  make dev-down   placeholder; teardown starts in a later task"
	@echo "  make api        placeholder; API server starts in a later task"
	@echo "  make worker     placeholder; worker starts in a later task"
	@echo "  make web        placeholder; web app starts in a later task"
	@echo "  make migrate    placeholder; migrations start in a later task"

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	gofmt -w .

vet:
	go vet ./...

check: fmt vet test lint

generate:
	go generate ./...

dev:
	@echo "dev environment not implemented yet; added in W1-03+"

dev-down:
	@echo "dev environment not implemented yet; added in W1-03+"

api:
	@echo "api not implemented yet"

worker:
	@echo "worker not implemented yet"

web:
	@echo "web not implemented yet"

migrate:
	@echo "migrate not implemented yet"
