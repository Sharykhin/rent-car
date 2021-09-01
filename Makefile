.PHONY: web test migrate-create migrate-up migrate-down

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

current_dir = $(shell pwd)
network ?= host

web:
	LOG_LEVEL=debug go run cmd/web/main.go

t:
	LOG_LEVEL=debug go run cmd/cli/main.go

test:
	go test -v ./... | { grep -v 'no test files'; true; }

migrate-create:
	docker run -v ${current_dir}/db/migrations:/db/migrations migrate/migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	docker run --network $(network) -v ${current_dir}/db/migrations:/db/migrations migrate/migrate -database ${POSTGRES_URL} -path db/migrations up

migrate-down:
	docker run --network $(network) -v ${current_dir}/db/migrations:/db/migrations migrate/migrate -database ${POSTGRES_URL} -path db/migrations down 1

mockery:
	docker run -w /app -v ${current_dir}:/app vektra/mockery --dir=$(dir) --all --output=$(output)