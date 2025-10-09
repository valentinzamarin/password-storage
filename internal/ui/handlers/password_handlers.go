package handlers

import "fyne.io/fyne/v2"

func CopyToClipboard(text string) {
	window := fyne.CurrentApp().Driver().AllWindows()[0]
	window.Clipboard().SetContent(text)
}
