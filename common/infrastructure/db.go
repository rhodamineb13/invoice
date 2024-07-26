package infrastructure

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func newDBConnection() *sqlx.DB {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/invoices", config.User, config.Password))
	if err != nil {
		panic(err)
	}
	db.PingContext(context.TODO())
	if err = migrateUp(db); err != nil {
		panic(err)
	}

	return db
}

func migrateUp(db *sqlx.DB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance("file://migration", config.DBName, driver)
	if err != nil {
		return err
	}

	err = mig.Up()
	if err != nil {
		return err
	}
	return nil
}
