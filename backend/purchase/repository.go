package purchase

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

// Repository represent the article's repository contract
type Repository interface {
	FindAll() (res []schema.Purchase, err error)
	Create(purchase schema.Purchase) (schema.Purchase, error)
}
