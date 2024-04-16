include .env
export
export GO111MODULE        ?= on

run:
	go run main.go serve-http

run-reload:
	find . | entr -cr $(MAKE) run

.PHONY: migration
migration:
	go run main.go db:migrate create $(name) sql

migration-status:
	go run main.go db:migrate status

migrate:
	go run main.go db:migrate up

migrate-back:
	go run main.go db:migrate down

migrate-reset:
	go run main.go db:migrate reset

generate-docs:
	`go env GOPATH`/bin/swag fmt
	`go env GOPATH`/bin/swag init -g internal/app/server/routes.go --parseDependency true --parseInternal true
