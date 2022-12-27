package user

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

// Repository represent the article's repository contract
type Repository interface {
	FindAll() (res []schema.UserRead, err error)
	Create(user schema.User) (schema.User, error)
	Get(id int) (schema.UserRead, error)
	Update(user schema.User) error
	Delete(id []int) error
}
