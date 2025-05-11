package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/itsDrac/wobot/store"
	_ "github.com/mattn/go-sqlite3"
)

const databaseName string = "database.db"

func main() {
	// Make DatabaseConfig.
	dbConfig := store.DatabaseConfig{
		DatabaseName:    databaseName,
		MigrationFolder: "./migrations/",
	}

	_ = store.NewSQLiteStore(&dbConfig)

}
