package infrastructure

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = newDBConnection()

func newDBConnection() *sqlx.DB {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		panic(err)
	}
	if err := migrateUp(db); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	return db
}

func migrateUp(db *sqlx.DB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		config.DBName,
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}

	return nil
}
