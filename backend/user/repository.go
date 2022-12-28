package user

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/user/filter"
)

// Repository represent the article's repository contract
type Repository interface {
	FindAll() (res []schema.UserRead, err error)
	FindByFilter(filter filter.UserFilter) ([]schema.UserRead, error)
	Create(user schema.User) (schema.User, error)
	Get(id int) (schema.UserRead, error)
	GetByEmail(email string) (schema.User, error)
	Update(user schema.User) error
	Delete(id []int) error
}
