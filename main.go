package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/itsDrac/wobot/handler"
	"github.com/itsDrac/wobot/service"
	"github.com/itsDrac/wobot/store"
	_ "github.com/mattn/go-sqlite3"
)

const databaseName string = "database.db"

func main() {
	// Make DatabaseConfig.
	var api API
	dbConfig := store.DatabaseConfig{
		DatabaseName:    databaseName,
		MigrationFolder: "./migrations/",
	}

	databaseStore := store.NewSQLiteStore(&dbConfig)
	defer databaseStore.DB.Close()
	serv := service.NewService(&databaseStore)
	router := handler.NewChiRouter()
	api = handler.ChiHandler{Service: &serv}
	api.Mount(router)
	api.Run(router)
}
