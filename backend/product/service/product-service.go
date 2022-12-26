package service

import (
	"github.com/kleberalves/problemCompanyApp/backend/product"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type productService struct {
	productRepo product.Repository
}

func NewProductService(repo product.Repository) product.Service {
	return &productService{
		productRepo: repo,
	}
}

func (srv *productService) FindAll() (res []schema.Product, err error) {
	return srv.productRepo.FindAll()
}

func (srv *productService) Create(input schema.Product) (schema.Product, error) {
	return srv.productRepo.Create(input)
}
