package service

import (
	"errors"

	"github.com/kleberalves/problemCompanyApp/backend/product"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type service struct {
	repo product.Repository
}

func NewProductService(repo product.Repository) product.Service {
	return &service{
		repo: repo,
	}
}

func (srv *service) FindAll() (res []schema.Product, err error) {
	return srv.repo.FindAll()
}

func (srv *service) Create(input schema.Product) (schema.Product, error) {
	return srv.repo.Create(input)
}

func (srv *service) Update(item schema.Product) error {

	if item.ID <= 0 {
		return errors.New("ID not found")
	}

	o, err := item.DeletedAt.Value()

	if o == nil && err == nil {
		srv.repo.UpdateDeletedAt(item)
	}

	return srv.repo.Update(item)
}

func (srv *service) Delete(itemIds []int) error {
	return srv.repo.Delete(itemIds)
}
