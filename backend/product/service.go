package product

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type Service interface {
	FindAll() (res []schema.Product, err error)
	Create(item schema.Product) (schema.Product, error)
	Update(item schema.Product) error
	Delete(itemIds []int) error
}
