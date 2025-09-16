package views

import (
	"password-storage/internal/ui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Searchbar struct{}

func NewSearchbarView() *Searchbar {
	return &Searchbar{}
}

func (h *Searchbar) Render() fyne.CanvasObject {
	/*
		next time
		ill use
		something
		with html + js
	*/
	searchEntry := components.CreateInputField("Search")

	searchEntry.OnChanged = func(text string) {

	}

	return container.NewVBox(
		searchEntry,
	)
}
