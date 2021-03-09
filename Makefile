.PHONY: web test migrate-create

current_dir = $(shell pwd)

web:
	go run cmd/web/main.go

test:
	go run cmd/cli/main.go

migrate-create:
	docker run -v ${current_dir}/db/migrations:/db/migrations migrate/migrate create -ext sql -dir db/migrations -seq $(name)