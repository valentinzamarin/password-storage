package entities

import (
	"testing"
)

func TestPassword_Validate(t *testing.T) {
	tests := []struct {
		name      string
		password  *Password
		wantError bool
		errMsg    string
	}{
		{
			name: "valid password",
			password: &Password{
				URL:      "https://example.com",
				Login:    "user",
				Password: "secret",
			},
			wantError: false,
		},
		{
			name: "missing URL",
			password: &Password{
				URL:      "",
				Login:    "user",
				Password: "secret",
			},
			wantError: true,
			errMsg:    "url is required",
		},
		{
			name: "missing login",
			password: &Password{
				URL:      "https://example.com",
				Login:    "",
				Password: "secret",
			},
			wantError: true,
			errMsg:    "login is required",
		},
		{
			name: "missing password",
			password: &Password{
				URL:      "https://example.com",
				Login:    "user",
				Password: "",
			},
			wantError: true,
			errMsg:    "password is required",
		},
		{
			name: "all fields missing",
			password: &Password{
				URL:      "",
				Login:    "",
				Password: "",
			},
			wantError: true,
			errMsg:    "url is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.password.Validate()

			if tt.wantError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				} else if tt.errMsg != "" && err.Error() != tt.errMsg {
					t.Errorf("Expected error message '%s', got '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestNewPassword(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		login       string
		password    string
		description string
		wantError   bool
	}{
		{
			name:        "create valid password",
			url:         "https://example.com",
			login:       "user",
			password:    "secret",
			description: "test account",
			wantError:   false,
		},
		{
			name:        "create with empty description",
			url:         "https://example.com",
			login:       "user",
			password:    "secret",
			description: "",
			wantError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPassword(tt.url, tt.login, tt.password, tt.description)

			if tt.wantError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if got.URL != tt.url {
					t.Errorf("Expected URL %s, got %s", tt.url, got.URL)
				}
				if got.Login != tt.login {
					t.Errorf("Expected Login %s, got %s", tt.login, got.Login)
				}
				if got.Password != tt.password {
					t.Errorf("Expected Password %s, got %s", tt.password, got.Password)
				}
				if got.Description != tt.description {
					t.Errorf("Expected Description %s, got %s", tt.description, got.Description)
				}
			}
		})
	}
}

func TestNewPassword_Validation(t *testing.T) {
	url := "https://example.com"
	login := "user"
	password := "secret"
	description := "test"

	p, err := NewPassword(url, login, password, description)
	if err != nil {
		t.Fatalf("NewPassword failed: %v", err)
	}

	if err := p.Validate(); err != nil {
		t.Errorf("NewPassword created invalid object: %v", err)
	}
}
