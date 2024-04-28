package cmd

import (
	"database/sql"
	"rr-backend/config/env"

	"github.com/pressly/goose/v3"
)

func RunMigration(db *sql.DB) {
	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	err := goose.Up(db, env.MigrationDir())
	if err != nil {
		panic(err)
	}
}
