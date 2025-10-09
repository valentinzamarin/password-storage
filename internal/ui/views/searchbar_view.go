package views

import (
	"fmt"
	"password-storage/internal/app/services"
	"password-storage/internal/ui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Searchbar struct {
	passwordService *services.PasswordService
}

func NewSearchbarView(passwordService *services.PasswordService) *Searchbar {
	return &Searchbar{
		passwordService: passwordService,
	}
}

func (h *Searchbar) Render() fyne.CanvasObject {
	/*
		next time
		ill use
		something
		with html + js
	*/
	searchEntry := components.CreateInputField("Search")

	resultsContainer := container.NewVBox()

	searchEntry.OnChanged = func(text string) {

		resultsContainer.Objects = nil

		if text == "" {

			resultsContainer.Refresh()
			return
		}

		passwords := h.passwordService.SearchPassword(text)

		for _, pwd := range passwords {

			label := widget.NewLabel(fmt.Sprintf("URL: %s", pwd.URL))
			resultsContainer.Add(label)
		}

		resultsContainer.Refresh()
	}

	return container.NewVBox(
		searchEntry,
		resultsContainer,
	)
}
