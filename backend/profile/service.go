package profile

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type Service interface {
	AddProfile(userId int, typo int) (schema.Profile, error)
	RemoveProfiles(userIds []int, typo int) error
	FindAll() ([]schema.Profile, error)
}
