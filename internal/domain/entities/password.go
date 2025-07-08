package entities

import "errors"

type Password struct {
	ID          int
	URL         string
	Login       string
	Password    string
	Description string
}

func (p *Password) Validate() error {
	if p.URL == "" {
		return errors.New("url is required")
	}
	if p.Login == "" {
		return errors.New("login is required")
	}
	if p.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func NewPassword(url, login, password, description string) (*Password, error) {
	p := &Password{
		URL:         url,
		Login:       login,
		Password:    password,
		Description: description,
	}

	return p, nil
}
