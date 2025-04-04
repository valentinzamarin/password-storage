package main

import (
	"database/sql"
)

type PasswordRecord struct {
	URL      string
	Login    string
	Password string
}

func GetPasswordsFromDB(db *sql.DB) ([]PasswordRecord, error) {
	query := `SELECT url, login, password FROM passwords`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwords []PasswordRecord
	for rows.Next() {
		var record PasswordRecord
		if err := rows.Scan(&record.URL, &record.Login, &record.Password); err != nil {
			return nil, err
		}
		passwords = append(passwords, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return passwords, nil
}
