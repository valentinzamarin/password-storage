package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"password-storage/internal/domain/entities"
	"password-storage/internal/domain/repositories"
)

type BasePasswordRepository struct {
	db *sql.DB
}

func NewBasePasswordRepository(db *sql.DB) repositories.PasswordRepo {
	return &BasePasswordRepository{db: db}
}

func (p *BasePasswordRepository) AddPassword(password entities.Password) error {
	query := `
		INSERT INTO passwords (url, login, password, description) 
		VALUES (?, ?, ?, ?)
	`

	result, err := p.db.Exec(query, password.URL, password.Login, password.Password, password.Description)
	if err != nil {
		return fmt.Errorf("failed to insert password: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were inserted")
	}

	return nil
}

func (p *BasePasswordRepository) GetAllPasswords() ([]*entities.Password, error) {
	query := `SELECT * FROM passwords`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var passwords []*entities.Password

	for rows.Next() {
		var pwd entities.Password
		err := rows.Scan(&pwd.ID, &pwd.URL, &pwd.Login, &pwd.Password, &pwd.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		passwords = append(passwords, &pwd)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return passwords, nil
}

func (p *BasePasswordRepository) DeletePasswordById(id int) error {

	query := "DELETE FROM passwords WHERE id = $1"

	result, err := p.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("exec err: %w", err)
	}

	_, err = result.RowsAffected()

	return nil
}
