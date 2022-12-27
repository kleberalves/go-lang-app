package purchase

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

// Repository represent the article's repository contract
type Repository interface {
	FindAll() (res []schema.Purchase, err error)
	GetByUser(userId int) ([]schema.Purchase, error)
	Create(purchase schema.Purchase) (schema.Purchase, error)
	Delete(purchaseIds []int) error
}
