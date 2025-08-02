package domainevents

const PasswordTopic = "passwords"

type AddedPasswordEvent struct {
	URL         string
	Login       string
	Password    string
	Description string
}

type RemovedPasswordEvent struct {
	ID uint
}

type UpdatePasswordEvent struct {
	ID          uint
	Description string
}
