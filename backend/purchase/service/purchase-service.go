package service

import (
	"time"

	"github.com/kleberalves/problemCompanyApp/backend/purchase"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
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

func (srv *service) GetByUser(userId int) ([]schema.Purchase, error) {
	//TODO: Check if authenticated user is owner or have a Salesman role
	return srv.repo.GetByUser(userId)
}

func (srv *service) Create(input schema.Purchase) (schema.Purchase, error) {
	input.PurchasedAt = time.Now()
	return srv.repo.Create(input)
}

func (srv *service) Delete(purchaseIds []int) error {
	return srv.repo.Delete(purchaseIds)
}
