package store

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DatabaseConfig struct {
	DatabaseName    string
	MigrationFolder string
}

func MigrateUp(con *DatabaseConfig, db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Println("Can not get the driver for migration.")
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", con.MigrationFolder),
		"sqlite3", driver)
	if err != nil {
		fmt.Println("Can not get migration instance.")
		return err
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No change in migration.")
			return nil
		}
		fmt.Println("Error in migration.")
		return err
	}
	return nil
}
