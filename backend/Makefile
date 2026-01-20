include app.env
.PHONY: all

up-db-container:
	docker run --name testDB -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.4-alpine3.20

createdb:
	docker exec -it testDB createdb alexdb

dropdb:
	docker exec -it testDB dropdb alexdb

migrate-files:
	migrate create -ext sql -dir db/migrations -seq $(NAME)

migrate-down:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/alexdb?sslmode=disable" -verbose down
	
migrate-up:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/alexdb?sslmode=disable" -verbose up

migrate-up-docker:
	docker exec -it alex-app-db-1 migrate -path /migrations -database "postgres://root:secret@testDB:5432/alexdb?sslmode=disable" -verbose up

sqlc:
	sqlc generate

server:
	air -c .air.toml

up-backend:
	docker compose up

# Apply migrations
migrate-apply:
	go run cmd/migrate/main.go -action=apply

# Check migration status
migrate-status:
	go run cmd/migrate/main.go -action=status

# Compute migration hash
migrate-hash:
	go run cmd/migrate/main.go -action=hash

migrate-diff:
	go run cmd/migrate/main.go -action=diff

# Run migrations for specific environment
migrate-apply-dev:
	go run cmd/migrate/main.go -action=apply -env=dev

migrate-status-dev:
	go run cmd/migrate/main.go -action=status -env=dev

# Build migration binary
build-migrate:
	go build -o bin/migrate cmd/migrate/main.go

# Run migration binary (if built)
migrate-bin:
	./bin/migrate -action=apply