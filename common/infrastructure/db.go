package infrastructure

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func newDBConnection() *sqlx.DB {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		panic(err)
	}
	db.PingContext(context.TODO())

	return db
}
