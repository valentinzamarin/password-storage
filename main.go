package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	passwordApp := app.New()

	desktopWindow := passwordApp.NewWindow("Password Storage")
	desktopWindow.Resize(fyne.NewSize(720, 480))

	db, err := InitializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	addPasswordForm := CreateForm(db)
	desktopWindow.SetContent(addPasswordForm)

	desktopWindow.ShowAndRun()
}
