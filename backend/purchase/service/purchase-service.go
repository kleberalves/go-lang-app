package service

import (
	"github.com/kleberalves/problemCompanyApp/backend/purchase"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type purchaseService struct {
	purchaseRepo purchase.Repository
}

func NewPurchaseService(repo purchase.Repository) purchase.Service {
	return &purchaseService{
		purchaseRepo: repo,
	}
}

func (srv *purchaseService) FindAll() (res []schema.Purchase, err error) {
	return srv.purchaseRepo.FindAll()
}

func (srv *purchaseService) Create(input schema.Purchase) (schema.Purchase, error) {
	return srv.purchaseRepo.Create(input)
}
