package profile

import (
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

// Repository represent the article's repository contract
type Repository interface {
	AddProfile(userId int, typ enums.TypeUser) (schema.Profile, error)
	RemoveProfiles(userIds []int, typ enums.TypeUser) error
	FindAll() ([]schema.Profile, error)
}
