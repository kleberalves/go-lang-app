package schema

import (
	"gorm.io/gorm"
)

// Hack to hide fields (ex Password)
// https://stackoverflow.com/questions/44003152/hide-fields-in-golang-gorm
type UserRead struct {
	gorm.Model
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"unique" `
	Profiles  []Profile `json:"profiles" gorm:"foreignkey:UserID"`
}
