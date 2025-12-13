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
