package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations(db *sql.DB, migrationsDir string) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (version VARCHAR(255) PRIMARY KEY);`); err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return fmt.Errorf("failed to query applied migrations: %w", err)
	}
	defer rows.Close()

	appliedMigrations := make(map[string]bool)
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return fmt.Errorf("failed to scan applied migration version: %w", err)
		}
		appliedMigrations[version] = true
	}

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("could not read migrations directory %s: %w", migrationsDir, err)
	}

	var migrationFiles []os.DirEntry
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".up.sql") && !file.IsDir() {
			migrationFiles = append(migrationFiles, file)
		}
	}

	sort.Slice(migrationFiles, func(i, j int) bool {
		return migrationFiles[i].Name() < migrationFiles[j].Name()
	})

	for _, file := range migrationFiles {
		fileName := file.Name()

		if _, ok := appliedMigrations[fileName]; ok {
			log.Printf("Migration %s already applied, skipping", fileName)
			continue
		}

		migrationSQL, err := os.ReadFile(filepath.Join(migrationsDir, fileName))
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", fileName, err)
		}

		sqlCommands := strings.Split(string(migrationSQL), ";")

		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction for %s: %w", fileName, err)
		}

		for _, command := range sqlCommands {
			command = strings.TrimSpace(command)
			if command == "" {
				continue
			}

			if _, err := tx.Exec(command); err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to apply migration %s, command: %s, error: %w", fileName, command, err)
			}
		}

		if _, err := tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)", fileName); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration version %s: %w", fileName, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction for %s: %w", fileName, err)
		}

	}

	return nil
}
