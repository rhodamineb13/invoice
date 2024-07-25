package infrastructure

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDBConnection() *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s %s", config.Host, config.Port))
	if err != nil {
		panic(err)
	}
	return db
}
