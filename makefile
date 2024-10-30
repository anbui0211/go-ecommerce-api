GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= "root:root1234@tcp(127.0.0.1:33306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema

# Tên của ứng dụng
APP_NAME = server

# Chạy ứng dụng
dev:
	go run ./cmd/$(APP_NAME)

make_build:
	docker compose up -d --build
	docker compose ps
docker_stop:
	docker compose down

docker_up:
	docker compose up -d

up_by_one:
	 @GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one

create_migration:
	@goose -dir=${GOOSE_MIGRATION_DIR} create $(name) sql
upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

# Generate code
sqlgen_v2:
	@docker pull sqlc/sqlc
	@docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc generate
sqlgen:
	docker pull sqlc/sqlc
	sqlc generate

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev upse downse resetse docker_build docker_up docker_stop swag

.PHONY: air

# Create a new sql file

