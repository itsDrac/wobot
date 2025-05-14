package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/itsDrac/wobot/docs"
	"github.com/itsDrac/wobot/handler"
	"github.com/itsDrac/wobot/service"
	"github.com/itsDrac/wobot/store"
	_ "github.com/mattn/go-sqlite3"
)

const databaseName string = "database.db"

// @title Wobot API
// @version 1.0
// @description This is a Wobot API, where users can upload files and get their storage information.

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @tokenUrl http://localhost:8080/api/v1/users/login

// @contact.name Sahaj
// @contact.email gpt.sahaj28@gmail.com
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
