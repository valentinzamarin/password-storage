package main

import (
	"log"
	"password-storage/internal/app/services"
	"password-storage/internal/infrastructure/sqlite"
	"password-storage/internal/ui"
)

func main() {

	basePath := "./passwords.db"

	db, err := sqlite.NewConnection(basePath)
	if err != nil {
		log.Fatalf("db conn error: %v", err)
	}
	defer db.Close()

	migrationsPath := "./migrations"
	if err := sqlite.RunMigrations(db, migrationsPath); err != nil {
		log.Fatalf("could not run migrations: %v", err)
	}

	passwordRepo := sqlite.NewBasePasswordRepository(db)
	passwordService := services.NewPasswordService(passwordRepo)

	uiApp := ui.NewApp(passwordService)

	uiApp.Run()
}
