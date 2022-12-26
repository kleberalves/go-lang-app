package schema

import (
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Type   enums.TypeUser `json:"type" gorm:"unique"`
	UserID int            `gorm:"unique"`
}
