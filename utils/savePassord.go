package utils

import (
	"database/sql"
	"fmt"
)

func SavePasswordToDB(db *sql.DB, url, login, password string) error {
	query := `INSERT INTO passwords (url, login, password) VALUES (?, ?, ?)`
	_, err := db.Exec(query, url, login, password)
	if err != nil {
		return fmt.Errorf("failed to insert data into database: %w", err)
	}
	return nil
}
