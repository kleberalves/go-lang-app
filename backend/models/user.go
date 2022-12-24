package models

import (
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Email     string         `json:"email"`
	Type      enums.TypeUser `json:"type"`
	Password  string         `json:"password"`
}
