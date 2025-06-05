package main

import (
	"log"

	"github.com/awe8128/backend/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config file", err)
	}
	migration, err := migrate.New(config.MigrationSource, config.DBSource)

	if err != nil {
		log.Fatal("Cannot create migration instance: ", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to run migration up: ", err)
	}
	log.Println("Migration completed successfully")
}
