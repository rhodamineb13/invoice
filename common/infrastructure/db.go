package infrastructure

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func newDBConnection() *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/test", config.Host, config.Port))
	if err != nil {
		panic(err)
	}
	return db
}
