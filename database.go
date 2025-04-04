package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./passwords.db")
	if err != nil {
		return nil, err
	}

	query := `
    CREATE TABLE IF NOT EXISTS passwords (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        url TEXT NOT NULL,
        login TEXT NOT NULL,
        password TEXT NOT NULL
    );`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
