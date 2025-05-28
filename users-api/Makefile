
MIGRATIONS_DIR=./database/migrations

# database settings
DATABASE_NAME=login_db
DATABASE_USER=login
DATABASE_PASSWORD=login
DATABASE_PORT=5432
DATABASE_HOST=localhost
DATABASE_URL=postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable

.PHONY: tests

tests:
	go test -v -cover -race ./tests/...

run:
	go run cmd/*.go 

# Migration download
migrate-download:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz & mv migrate /usr/local/bin/

# Mockgen download
mockgen-download:
	sudo apt install mockgen 

# Linter download
linter-download:
	curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v1.46.2/golangci-lint-1.46.2-linux-amd64.tar.gz | tar -xzv & sudo mv golangci-lint-1.46.2-linux-amd64/golangci-lint /usr/local/bin/

linter-run:
	golangci-lint run 


create-migrations: 
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" up 
	
migrate-last-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down 1
