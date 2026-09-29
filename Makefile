.PHONY: help test lint fmt vet check generate dev dev-down dev-status dev-logs api worker web migrate openapi-check

COMPOSE := docker compose --project-name runphase -f deploy/compose/docker-compose.yml

help:
	@echo "Runphase"
	@echo "  make help       list targets"
	@echo "  make test       run go test ./..."
	@echo "  make lint       run golangci-lint"
	@echo "  make fmt        rewrite Go sources with gofmt -w"
	@echo "  make vet        run go vet ./..."
	@echo "  make check      run fmt, vet, test, and lint"
	@echo "  make generate   run go generate ./..."
	@echo "  make dev        start local Postgres and Temporal"
	@echo "  make dev-status show local infrastructure status"
	@echo "  make dev-logs   show recent local infrastructure logs"
	@echo "  make dev-down   stop local Postgres and Temporal; keep data"
	@echo "  make api        run the API server locally"
	@echo "  make worker     placeholder; worker starts in a later task"
	@echo "  make web        placeholder; web app starts in a later task"
	@echo "  make migrate    apply Runphase database migrations"
	@echo "  make openapi-check  validate the OpenAPI document"

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
	$(COMPOSE) up -d --wait

dev-down:
	$(COMPOSE) down

dev-status:
	$(COMPOSE) ps

dev-logs:
	$(COMPOSE) logs --tail=200

api:
	go run ./cmd/apiserver

worker:
	@echo "worker not implemented yet"

web:
	@echo "web not implemented yet"

migrate:
	go run ./cmd/migrate

openapi-check:
	go test ./api/openapi/
