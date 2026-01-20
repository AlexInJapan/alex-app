package main

import (
	"context"
	"log"

	"github.com/awe8128/backend/api"
	db "github.com/awe8128/backend/db/sqlc"
	"github.com/awe8128/backend/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	log.Println("Starting Migration")
	migration, err := migrate.New(config.MigrationSource, config.DBSource)

	if err != nil {
		log.Fatal("Cannot create migration instance: ", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to run migration up: ", err)
	}
	log.Println("Migration completed successfully")

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
