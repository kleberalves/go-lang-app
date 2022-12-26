package purchase

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type Service interface {
	FindAll() (res []schema.Purchase, err error)
	Create(purchase schema.Purchase) (schema.Purchase, error)
}
