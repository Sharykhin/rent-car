.PHONY: web test migrate-create migrate-up migrate-down

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

current_dir = $(shell pwd)

web:
	LOG_LEVEL=debug go run cmd/web/main.go

test:
	LOG_LEVEL=debug go run cmd/cli/main.go

migrate-create:
	docker run -v ${current_dir}/db/migrations:/db/migrations migrate/migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	docker run --network host -v ${current_dir}/db/migrations:/db/migrations migrate/migrate -database ${POSTGRES_URL} -path db/migrations up

migrate-down:
	docker run --network host -v ${current_dir}/db/migrations:/db/migrations migrate/migrate -database ${POSTGRES_URL} -path db/migrations down 1