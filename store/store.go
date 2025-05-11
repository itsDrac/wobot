package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Store interface {
}

type SQLiteStore struct {
	DB *sql.DB
}

func NewSQLiteStore(con *DatabaseConfig) SQLiteStore {
	db, err := sql.Open("sqlite3", con.DatabaseName)
	if err != nil {
		fmt.Errorf("Error in opening the database. %s", err.Error())
	}
	// TODO: Maybe make this into a go routine.
	if err := MigrateUp(con, db); err != nil {
		fmt.Errorf("Error in migrating. %s", err.Error())
	}
	return SQLiteStore{
		DB: db,
	}
}
