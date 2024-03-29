package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique" `
	Password  string
	Profiles  []Profile `gorm:"foreignkey:UserID"`
}
