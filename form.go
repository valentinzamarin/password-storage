package main

import (
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/valentinzamarin/password-storage/utils"
)

func CreateForm(db *sql.DB) fyne.CanvasObject {
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("URL (example.com)")

	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Login")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()
	errorLabel.TextStyle.Bold = true
	errorLabel.Importance = widget.HighImportance

	passwordList := widget.NewLabel("No passwords saved yet.")
	passwordList.Wrapping = fyne.TextWrapWord

	updatePasswordList := func() {
		passwords, err := GetPasswordsFromDB(db)
		if err != nil {
			passwordList.SetText(fmt.Sprintf("Error loading passwords: %v", err))
			return
		}

		if len(passwords) == 0 {
			passwordList.SetText("No passwords saved yet.")
			return
		}

		listText := ""
		for _, p := range passwords {
			listText += fmt.Sprintf("URL: %s, Login: %s, Password: %s\n", p.URL, p.Login, p.Password)
		}
		passwordList.SetText(listText)
	}

	updatePasswordList()

	addButton := widget.NewButton("Add", func() {
		if urlEntry.Text == "" || loginEntry.Text == "" || passwordEntry.Text == "" {
			errorLabel.SetText("Error: All fields must be filled!")
			errorLabel.Show()
			return
		}

		err := utils.SavePasswordToDB(db, urlEntry.Text, loginEntry.Text, passwordEntry.Text)
		if err != nil {
			errorLabel.SetText(fmt.Sprintf("Error saving to database: %v", err))
			errorLabel.Show()
			return
		}

		errorLabel.Hide()
		fmt.Println("Added:", urlEntry.Text, loginEntry.Text, passwordEntry.Text)
		urlEntry.SetText("")
		loginEntry.SetText("")
		passwordEntry.SetText("")

		updatePasswordList()
	})

	form := container.NewVBox(
		widget.NewLabel("Add password:"),
		urlEntry,
		loginEntry,
		passwordEntry,
		errorLabel,
		addButton,
	)

	content := container.NewVBox(
		form,
		widget.NewLabel("Saved passwords:"),
		passwordList,
	)

	return content
}
