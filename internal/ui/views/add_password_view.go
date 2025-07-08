package views

import (
	"password-storage/internal/app/services"
	"password-storage/internal/domain/entities"
	"password-storage/internal/ui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type AddPasswordView struct {
	passwordService *services.PasswordService
	window          fyne.Window
	inputURL        *widget.Entry
	inputLogin      *widget.Entry
	inputPassword   *widget.Entry
}

func NewAddPasswordView(passwordService *services.PasswordService, window fyne.Window) *AddPasswordView {
	return &AddPasswordView{
		passwordService: passwordService,
		window:          window,
		inputURL:        components.CreateInputField("URL"),
		inputLogin:      components.CreateInputField("Login"),
		inputPassword:   components.CreateInputField("Password"),
	}
}

func (v *AddPasswordView) Render() fyne.CanvasObject {
	submitButton := widget.NewButton("Add password", v.handleSubmit)

	return container.NewVBox(
		v.inputURL,
		v.inputLogin,
		v.inputPassword,
		submitButton,
	)
}

func (v *AddPasswordView) handleSubmit() {
	url := v.inputURL.Text
	login := v.inputLogin.Text
	password := v.inputPassword.Text
	description := ""

	newPassword, err := entities.NewPassword(url, login, password, description)
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}

	validationErr := newPassword.Validate()
	if validationErr != nil {
		dialog.ShowError(validationErr, v.window)
		return
	}

	serviceErr := v.passwordService.AddNewPassword(url, login, password, "")
	if serviceErr != nil {
		dialog.ShowError(serviceErr, v.window)
		return
	}

	dialog.ShowInformation("success", "Password added", v.window)
	v.clearForm()
}

func (v *AddPasswordView) clearForm() {
	v.inputURL.SetText("")
	v.inputLogin.SetText("")
	v.inputPassword.SetText("")
}
