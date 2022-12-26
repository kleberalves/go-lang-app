package product

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type Service interface {
	FindAll() (res []schema.Product, err error)
	Create(user schema.Product) (schema.Product, error)
}
