package user

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type Service interface {
	FindAll() (res []schema.UserRead, err error)
	Create(user schema.User) (schema.User, error)
	Get(id int) (schema.UserRead, error)
	Update(user schema.User) error
	Delete(id []int) error

	AddProfile(userId int, typo int) (schema.Profile, error)
	DeleteProfile(userId int, typo int) error
}
