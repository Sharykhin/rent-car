.PHONY: start

start:
	go run cmd/web/main.go

test:
	go run cmd/cli/main.go