package ui

import (
	"password-storage/internal/app/services"
	"password-storage/internal/ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type App struct {
	fyneApp         fyne.App
	window          fyne.Window
	passwordService *services.PasswordService
}

func NewApp(passwordService *services.PasswordService) *App {
	a := app.New()
	w := a.NewWindow("Password Storage")

	return &App{
		fyneApp:         a,
		window:          w,
		passwordService: passwordService,
	}
}

func (a *App) Run() {
	a.window.Resize(fyne.NewSize(720, 600))
	a.window.SetContent(a.makeMainView())
	a.window.ShowAndRun()
}

func (a *App) makeMainView() fyne.CanvasObject {
	passwordListView := views.NewPasswordListView(a.passwordService, a.window)
	addPasswordView := views.NewAddPasswordView(a.passwordService, a.window)

	tabs := container.NewAppTabs(
		container.NewTabItem("Passwords", passwordListView.Render()),
		container.NewTabItem("Add", addPasswordView.Render()),
	)

	return tabs
}
