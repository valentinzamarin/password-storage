package ui

import (
	"fmt"
	"log"
	"password-storage/internal/app/encrypt"
	"password-storage/internal/app/events"
	"password-storage/internal/app/services"
	"password-storage/internal/ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	fyneApp         fyne.App
	window          fyne.Window
	passwordService *services.PasswordService
	eventBus        *events.EventBus
	encryptService  *encrypt.PasswordEncrypt
	authService     *services.AuthService
}

func NewApp(
	passwordService *services.PasswordService,
	eventBus *events.EventBus,
	encryptService *encrypt.PasswordEncrypt,
	authService *services.AuthService,
) *App {
	a := app.New()
	w := a.NewWindow("Password Storage")

	return &App{
		fyneApp:         a,
		window:          w,
		passwordService: passwordService,
		eventBus:        eventBus,
		encryptService:  encryptService,
		authService:     authService,
	}
}

func (a *App) Run() {
	a.window.Resize(fyne.NewSize(720, 600))
	a.showLoginOrSetupDialog()
	a.window.ShowAndRun()
}

func (a *App) showLoginOrSetupDialog() {
	isSet, err := a.authService.IsMasterPasswordSet()
	if err != nil {
		dialog.ShowError(fmt.Errorf("fatal database error: %w", err), a.window)
		a.window.Close()
		return
	}

	if isSet {
		a.showLoginDialog()
	} else {
		a.showSetupDialog()
	}
}

func (a *App) showLoginDialog() {
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter master password...")

	formDialog := dialog.NewForm("Unlock", "Unlock", "Cancel", []*widget.FormItem{
		widget.NewFormItem("Master Password", passwordEntry),
	}, func(confirmed bool) {
		if !confirmed {
			a.window.Close()
			return
		}

		err := a.authService.Authenticate(passwordEntry.Text)
		if err != nil {
			log.Println("Authentication failed:", err)
			dialog.ShowError(err, a.window)
			a.showLoginDialog()
			return
		}

		log.Println("Master password accepted. Loading main view.")
		a.loadMainView()
	}, a.window)

	formDialog.Resize(fyne.NewSize(400, 150))
	formDialog.Show()
}

func (a *App) showSetupDialog() {
	passwordEntry := widget.NewPasswordEntry()
	confirmEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Create a strong master password...")
	confirmEntry.SetPlaceHolder("Confirm password...")

	items := []*widget.FormItem{
		widget.NewFormItem("Master Password", passwordEntry),
		widget.NewFormItem("Confirm Password", confirmEntry),
	}

	formDialog := dialog.NewForm("Setup", "Create", "Cancel", items, func(confirmed bool) {
		if !confirmed {
			a.window.Close()
			return
		}

		if passwordEntry.Text == "" {
			dialog.ShowError(fmt.Errorf("password cannot be empty"), a.window)
			a.showSetupDialog()
			return
		}
		if passwordEntry.Text != confirmEntry.Text {
			dialog.ShowError(fmt.Errorf("passwords do not match"), a.window)
			a.showSetupDialog()
			return
		}

		err := a.authService.CreateMasterPassword(passwordEntry.Text)
		if err != nil {
			log.Println("Failed to create master password:", err)
			dialog.ShowError(err, a.window)
			a.showSetupDialog()
			return
		}

		log.Println("Master password created. Loading main view.")
		a.loadMainView()
	}, a.window)

	formDialog.Resize(fyne.NewSize(450, 200))
	formDialog.Show()
}

func (a *App) loadMainView() {
	mainContent := a.makeMainView()
	a.window.SetContent(mainContent)
}

func (a *App) makeMainView() fyne.CanvasObject {
	passwordListView := views.NewPasswordListView(a.passwordService, a.window, a.eventBus)
	addPasswordView := views.NewAddPasswordView(a.passwordService, a.window)

	tabs := container.NewAppTabs(
		container.NewTabItem("Passwords", passwordListView.Render()),
		container.NewTabItem("Add", addPasswordView.Render()),
	)

	return tabs
}
