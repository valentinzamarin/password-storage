package views

import (
	"password-storage/internal/app/events"
	"password-storage/internal/app/services"
	"password-storage/internal/domain/entities"
	domainevents "password-storage/internal/domain/events"
	"password-storage/internal/ui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

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
		// passwordItem := components.PasswordItem(password)
		// v.content.Add(passwordItem)

		id := password.ID
		passwordItem := components.PasswordItem(
			password,

			// onDelete
			func() {
				dialog.ShowConfirm("Delete password", "Are you sure?", func(ok bool) {
					if !ok {
						return
					}
					v.passwordService.DeletePassword(id)

				}, v.window)
			},

			// onUpdateDescription
			func(newDesc string) {
				err := v.passwordService.UpdatePassword(id, newDesc)
				if err != nil {
					dialog.ShowError(err, v.window)
				} else {
					dialog.ShowInformation("Saved", "Description updated", v.window)
				}
			},
		)
		v.content.Add(passwordItem)
	}
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
