package sqlite

type PasswordModel struct {
	ID          uint   `gorm:"primaryKey"`
	URL         string `gorm:"not null"`
	Login       string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Description string `gorm:""`
}

/*
GORM generates table names
table becomes "passwords_model
to avoid this - rename it.
*/
func (PasswordModel) TableName() string {
	return "passwords"
}
