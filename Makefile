include .env
export $(shell sed 's/=.*//' .env)

# Database URL
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

.PHONY: run db-create migrate-create migrate-up migrate-down

run:
	@go run cmd/api/main.go

db-create:
	@createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)

migrate-create:
	@migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	@migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	@migrate -path migrations -database "$(DB_URL)" down

tidy:
	@go mod tidy && go mod vendor
