package db

import (
	constant "rr-backend/lib/const"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

func GetSqlxClient() *sqlx.DB {
	db, err := sqlx.Connect("mysql", constant.DatabaseURL)
	if err != nil {
		panic(err)
	}

	return db
}
