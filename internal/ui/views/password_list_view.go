package views

import (
	"password-storage/internal/app/services"
	"password-storage/internal/domain/entities"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type TappableLabel struct {
	widget.Label
	onTapped func(string)
	window   fyne.Window
}

func NewTappableLabel(text string, window fyne.Window, onTapped func(string)) *TappableLabel {
	l := &TappableLabel{
		onTapped: onTapped,
		window:   window,
	}
	l.ExtendBaseWidget(l)
	l.SetText(text)
	return l
}

func (l *TappableLabel) Tapped(event *fyne.PointEvent) {
	if l.onTapped != nil {
		l.onTapped(l.Text)
		l.window.Clipboard().SetContent(l.Text)
		dialog.ShowInformation("", "Copied", l.window)
	}
}

type PasswordListView struct {
	passwordService *services.PasswordService
	window          fyne.Window
	list            *widget.List
	passwords       []*entities.Password
}

func NewPasswordListView(passwordService *services.PasswordService, window fyne.Window) *PasswordListView {
	view := &PasswordListView{
		passwordService: passwordService,
		window:          window,
		passwords:       []*entities.Password{},
	}

	view.createList()
	view.loadPasswords()

	return view
}

func (v *PasswordListView) createList() {
	v.list = widget.NewList(
		func() int {
			return len(v.passwords)
		},

		func() fyne.CanvasObject {
			urlLabel := NewTappableLabel("", v.window, nil)
			loginLabel := NewTappableLabel("", v.window, nil)
			passwordLabel := NewTappableLabel("", v.window, nil)

			deleteButton := widget.NewButton("Del", nil)

			return container.NewHBox(
				widget.NewLabel("URL:"),
				urlLabel,
				widget.NewLabel("Login:"),
				loginLabel,
				widget.NewLabel("Password:"),
				passwordLabel,
				deleteButton,
			)
		},

		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id >= len(v.passwords) {
				return
			}

			password := v.passwords[id]
			contentContainer := obj.(*fyne.Container)

			urlLabel := contentContainer.Objects[1].(*TappableLabel)
			loginLabel := contentContainer.Objects[3].(*TappableLabel)
			passwordLabel := contentContainer.Objects[5].(*TappableLabel)
			deleteButton := contentContainer.Objects[6].(*widget.Button)

			urlLabel.SetText(password.URL)
			loginLabel.SetText(password.Login)
			passwordLabel.SetText(password.Password)

			deleteButton.OnTapped = func() {
				v.passwordService.DeletePassword(password.ID)
			}

			urlLabel.Refresh()
			loginLabel.Refresh()
			passwordLabel.Refresh()
			deleteButton.Refresh()
		},
	)
}

func (v *PasswordListView) loadPasswords() {
	passwords, err := v.passwordService.GetPasswords()
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}

	v.passwords = passwords
	v.list.Refresh()
}

func (v *PasswordListView) Render() fyne.CanvasObject {
	return container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		v.list,
	)
}

func (v *PasswordListView) RefreshList() {
	v.loadPasswords()
}
