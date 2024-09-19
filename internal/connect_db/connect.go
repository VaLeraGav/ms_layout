package connect_db

import (
	"database/sql"
	"fmt"
)

func New(dbURL string, driver string) (*sql.DB, error) {
	const op = "connect_db.New"

	db, err := sql.Open(driver, dbURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}
