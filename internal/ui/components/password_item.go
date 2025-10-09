package components

import (
	"password-storage/internal/domain/entities"
	"password-storage/internal/ui/handlers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PasswordItem(password *entities.Password, onDelete func(), onUpdateDescription func(newDesc string)) fyne.CanvasObject {
	createClickableField := func(labelText, copyText string) fyne.CanvasObject {
		btn := widget.NewButton(labelText, func() {
			handlers.CopyToClipboard(copyText)
		})
		return btn
	}

	deleteBtn := widget.NewButton("❌", func() {
		onDelete()
	})

	var toggleButton *widget.Button
	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.SetText(password.Description)
	descriptionEntry.Wrapping = fyne.TextWrapWord

	saveButton := widget.NewButton("💾", func() {
		onUpdateDescription(descriptionEntry.Text)
	})

	descriptionContainer := container.NewVBox(
		widget.NewSeparator(),
		descriptionEntry,
		container.NewHBox(saveButton),
	)
	descriptionContainer.Hide()

	toggleButton = widget.NewButton("📝", func() {
		if descriptionContainer.Visible() {
			descriptionContainer.Hide()
			toggleButton.SetText("📝")
		} else {
			descriptionContainer.Show()
			toggleButton.SetText("🔼")
		}
	})

	urlField := createClickableField("URL: "+password.URL, password.URL)
	loginField := createClickableField("Login: "+password.Login, password.Login)
	passwordField := createClickableField("Password: "+password.Password, password.Password)

	topRow := container.NewHBox(
		urlField,
		loginField,
		passwordField,
		deleteBtn,
		toggleButton,
	)

	borderContainer := container.NewBorder(
		topRow,               // top
		descriptionContainer, // bottom
		nil, nil,             // left, right
		nil, // center
	)

	return borderContainer
}
