package views

import (
	"password-storage/internal/app/events"
	"password-storage/internal/app/services"
	"password-storage/internal/domain/entities"
	domainevents "password-storage/internal/domain/events"
	"password-storage/internal/ui/components"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type PasswordListView struct {
	passwordService   *services.PasswordService
	window            fyne.Window
	content           *fyne.Container
	eventBus          *events.EventBus
	allPasswords      []*entities.Password
	filteredPasswords []*entities.Password
	searchEntry       *widget.Entry
}

func NewPasswordListView(passwordService *services.PasswordService, window fyne.Window, eventBus *events.EventBus) *PasswordListView {

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search...")

	view := &PasswordListView{
		passwordService:   passwordService,
		window:            window,
		eventBus:          eventBus,
		content:           container.NewVBox(),
		allPasswords:      []*entities.Password{},
		filteredPasswords: []*entities.Password{},
		searchEntry:       searchEntry,
	}

	searchEntry.OnChanged = view.onSearch

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

	for _, password := range v.filteredPasswords {
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
			// only description ??
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

	v.allPasswords = passwords
	v.onSearch(v.searchEntry.Text)
}

// func (v *PasswordListView) Render() fyne.CanvasObject {
// 	/*
// 		no need in "NewBorder"
// 	*/
// 	return container.NewScroll(v.content)
// }

func (v *PasswordListView) Render() fyne.CanvasObject {
	return container.NewBorder(
		container.NewPadded(v.searchEntry),
		nil, nil, nil,
		container.NewScroll(v.content),
	)
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

func (v *PasswordListView) onSearch(query string) {
	query = strings.ToLower(strings.TrimSpace(query))

	v.filteredPasswords = nil

	if query == "" {
		v.filteredPasswords = v.allPasswords
	} else {
		for _, pwd := range v.allPasswords {
			if strings.Contains(strings.ToLower(pwd.URL), query) ||
				strings.Contains(strings.ToLower(pwd.Login), query) {
				v.filteredPasswords = append(v.filteredPasswords, pwd)
			}
		}
	}

	v.refreshContent()
}
