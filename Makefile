include app.env
.PHONY: all

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root alexdb
dropdb:
	docker exec -it postgres16 dropdb alexdb
migrate-up:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/alexdb?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/alexdb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run ./backend/cmd/server