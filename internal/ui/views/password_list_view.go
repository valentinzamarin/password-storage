package views

import (
	"password-storage/internal/app/events"
	"password-storage/internal/app/services"
	"password-storage/internal/domain/entities"
	domainevents "password-storage/internal/domain/events"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type TappableLabel struct {
	widget.Label
	window fyne.Window
}

func NewTappableLabel(text string, window fyne.Window) *TappableLabel {
	l := &TappableLabel{
		window: window,
	}
	l.ExtendBaseWidget(l)
	l.SetText(text)
	return l
}

func (l *TappableLabel) Tapped(event *fyne.PointEvent) {

	l.window.Clipboard().SetContent(l.Text)
	dialog.ShowInformation("Copied", "to the clipboard", l.window)
}

type PasswordListView struct {
	passwordService *services.PasswordService
	window          fyne.Window
	content         *fyne.Container
	passwords       []*entities.Password
	eventBus        *events.EventBus
}

func NewPasswordListView(passwordService *services.PasswordService, window fyne.Window, eventBus *events.EventBus) *PasswordListView {
	view := &PasswordListView{
		passwordService: passwordService,
		window:          window,
		passwords:       []*entities.Password{},
		eventBus:        eventBus,
		content:         container.NewVBox(),
	}

	view.createList()
	view.loadPasswords()
	view.subscribeToEvents()
	return view
}

func (v *PasswordListView) createList() {

	v.content = container.NewVBox()

	v.refreshContent()
}

func (v *PasswordListView) refreshContent() {
	v.content.RemoveAll()

	for _, password := range v.passwords {
		item := v.createPasswordItem(password)
		v.content.Add(item)
	}
}

func (v *PasswordListView) createPasswordItem(password *entities.Password) fyne.CanvasObject {

	urlLabel := NewTappableLabel(password.URL, v.window)
	loginLabel := NewTappableLabel(password.Login, v.window)
	passwordLabel := NewTappableLabel(password.Password, v.window)

	deleteButton := widget.NewButton("❌", func() {
		v.passwordService.DeletePassword(password.ID)
	})

	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.SetText(password.Description)
	descriptionEntry.Wrapping = fyne.TextWrapWord

	saveButton := widget.NewButton("💾", func() {

		password.Description = descriptionEntry.Text
		err := v.passwordService.UpdatePassword(password.ID, password.Description)
		if err != nil {
			dialog.ShowError(err, v.window)
		} else {
			dialog.ShowInformation("Changed", "", v.window)
		}
	})

	descriptionContainer := container.NewVBox(
		widget.NewSeparator(),

		descriptionEntry,
		saveButton,
	)
	descriptionContainer.Hide()

	var toggleButton *widget.Button
	toggleButton = widget.NewButton("📝", func() {
		if descriptionContainer.Visible() {
			descriptionContainer.Hide()
			toggleButton.SetText("📝")
		} else {
			descriptionContainer.Show()
			toggleButton.SetText("🔼")
		}
	})

	mainRowWithToggle := container.NewHBox(
		widget.NewLabel("URL:"),
		urlLabel,
		widget.NewLabel("Login:"),
		loginLabel,
		widget.NewLabel("Password:"),
		passwordLabel,
		deleteButton,
		toggleButton,
	)

	borderContainer := container.NewBorder(
		mainRowWithToggle,    // top
		descriptionContainer, // bottom
		nil,                  // left
		nil,                  // right
		nil,                  // center
	)

	return borderContainer
}

func (v *PasswordListView) loadPasswords() {
	passwords, err := v.passwordService.GetPasswords()
	if err != nil {
		dialog.ShowError(err, v.window)
		return
	}

	v.passwords = passwords
	v.refreshContent()
}

func (v *PasswordListView) Render() fyne.CanvasObject {
	/*
		no need in "NewBorder"
	*/
	return container.NewScroll(v.content)
}
func (v *PasswordListView) subscribeToEvents() {
	ch := v.eventBus.Subscribe(domainevents.PasswordTopic)

	go func() {
		for event := range ch {
			fyne.Do(func() {
				switch event.(type) {
				case domainevents.AddedPasswordEvent:
					v.loadPasswords()
				case domainevents.RemovedPasswordEvent:
					v.loadPasswords()
				}
			})
		}
	}()
}
