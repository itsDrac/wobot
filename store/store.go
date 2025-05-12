package store

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Store interface {
	CreateUser(context.Context, *User) error
	GetUserByUsername(context.Context, *User) error
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
