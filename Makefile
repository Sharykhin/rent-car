.PHONY: web test

web:
	go run cmd/web/main.go

test:
	go run cmd/cli/main.go