package product

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

// Repository represent the article's repository contract
type Repository interface {
	FindAll() (res []schema.Product, err error)
	Create(user schema.Product) (schema.Product, error)
}
