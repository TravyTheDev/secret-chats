package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func NewSqlStorage() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "db/app.db")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	return db, nil
}
