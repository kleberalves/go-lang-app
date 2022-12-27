package product

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

// Repository represent the article's repository contract
type Repository interface {
	FindAll() (res []schema.Product, err error)
	Create(item schema.Product) (schema.Product, error)
	Update(item schema.Product) error
	UpdateDeletedAt(item schema.Product) error
	Delete(itemIds []int) error
}
