package schema

import (
	"gorm.io/gorm"
)

// Hack to hide fields like password (sensitive) instead show empty
type UserRead struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string    `gorm:"unique" `
	Profiles  []Profile `gorm:"foreignkey:UserID"`
}
