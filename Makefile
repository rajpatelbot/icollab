# Load environment variables
include .env
export $(shell sed 's/=.*//' .env)

# Directories and DB connection
MIGRATIONS_DIR=migrations
DB_URL=postgres://$(USER):$(PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

# 🏗️ Build the Go binary
build:
	@go build -o bin/go_backend_setup main.go

# 🚀 Run the Go server
run:
	@go run cmd/server/main.go

# 🧱 Create a new empty migration file
migrate-create:
	@migrate create -ext sql -dir internal/migrations/$(subdir) -seq $(name)

# Apply migrations for a specific module
migrate-up:
	@migrate -path internal/migrations/$(subdir) -database $(DB_URL) up

# Rollback the last migration for a specific module
migrate-down:
	@migrate -path internal/migrations/$(subdir) -database $(DB_URL) down 1
