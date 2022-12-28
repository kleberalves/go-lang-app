package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/purchase"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services/security"
)

type service struct {
	repo purchase.Repository
}

func NewPurchaseService(repo purchase.Repository) purchase.Service {
	return &service{
		repo: repo,
	}
}

func (srv *service) FindAll() (res []schema.Purchase, err error) {
	return srv.repo.FindAll()
}

func (srv *service) GetByMyId(c *gin.Context) ([]schema.Purchase, error) {

	userId, err := security.ExtractTokenID(c)
	if err != nil {
		return []schema.Purchase{}, err
	}

	//TODO: Check if authenticated user is owner or have a Salesman role
	return srv.repo.GetByUser(userId)
}

func (srv *service) CreateByAuthenticatedSalesman(input schema.Purchase, c *gin.Context) (schema.Purchase, error) {

	userId, err := security.ExtractTokenID(c)
	if err != nil {
		return schema.Purchase{}, err
	}

	input.CodSalesman.Scan(userId)
	input.PurchasedAt = time.Now()
	return srv.repo.Create(input)
}

func (srv *service) CreateMyPurchase(input schema.Purchase, c *gin.Context) (schema.Purchase, error) {

	userId, err := security.ExtractTokenID(c)
	if err != nil {
		return schema.Purchase{}, err
	}
	input.CodCustomer = userId
	input.PurchasedAt = time.Now()
	return srv.repo.Create(input)
}

func (srv *service) Delete(purchaseIds []int) error {
	return srv.repo.Delete(purchaseIds)
}
