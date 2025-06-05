include app.env
.PHONY: all

createdb-local:
	docker exec -it postgres16 createdb alexdb
dropdb:
	docker exec -it postgres16 dropdb alexdb
migrate-up:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/alexdb?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/alexdb?sslmode=disable" -verbose down
migrate-up-docker:
	docker exec -it alex-app-db-1 migrate -path /migrations -database "postgres://root:secret@localhost:5432/alexdb?sslmode=disable" -verbose up

sqlc:
	sqlc generate

server:
	go run ./cmd/server

up-backend:
	docker compose up