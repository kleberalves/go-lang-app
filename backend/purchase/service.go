package purchase

import (
	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type Service interface {
	FindAll() (res []schema.Purchase, err error)
	GetByMyId(c *gin.Context) ([]schema.Purchase, error)
	CreateByAuthenticatedSalesman(purchase schema.Purchase, c *gin.Context) (schema.Purchase, error)
	CreateMyPurchase(purchase schema.Purchase, c *gin.Context) (schema.Purchase, error)
	Delete(purchaseIds []int) error
}
