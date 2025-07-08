package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewConnection(baseName string) (*sql.DB, error) {
	con, err := sql.Open("sqlite3", baseName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := con.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return con, nil
}
