package main

import (
	"log"
	"password-storage/internal/app/encrypt"
	"password-storage/internal/app/events"
	"password-storage/internal/app/services"

	"password-storage/internal/infrastructure/sqlite"
	"password-storage/internal/infrastructure/sqlite/auth"
	passwords "password-storage/internal/infrastructure/sqlite/password"

	"password-storage/internal/ui"
)

func main() {

	basePath := "./notebook.db"

	db, err := sqlite.NewConnection(basePath)
	if err != nil {
		log.Fatalf("db conn error: %v", err)
	}

	sqlite.Migrate(db)

	/* additional functional */
	eventBus := events.NewEventBus()
	encrypt := encrypt.NewPasswordEncrypt()

	/* db repos */
	authRepo := auth.NewAuthRepo(db)
	passwordRepo := passwords.NewGormPasswordRepository(db, encrypt)

	/* app services */
	authService := services.NewAuthService(authRepo, encrypt)
	passwordService := services.NewPasswordService(passwordRepo, eventBus, encrypt)

	uiApp := ui.NewApp(passwordService, eventBus, encrypt, authService)

	uiApp.Run()
}
