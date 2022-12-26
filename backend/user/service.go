package user

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type Service interface {
	FindAll() (res []schema.UserRead, err error)
	Create(user schema.User) (schema.User, error)
	AssociateProfile(userId int, typo int) schema.Profile
	RemoveProfile(userId int, typo int)
}
