package repository

import (
	"github.com/kleberalves/problemCompanyApp/backend/product"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"gorm.io/gorm"
)

type productRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(Conn *gorm.DB) product.Repository {
	return &productRepository{Conn}
}

func (repo *productRepository) FindAll() (res []schema.Product, err error) {

	var products []schema.Product
	errExec := repo.Conn.Model(&schema.Product{}).Find(&products).Error

	if errExec != nil {
		panic("Failed to retrieve all Products: " + err.Error())
	}

	return products, errExec
}

func (repo *productRepository) Create(input schema.Product) (schema.Product, error) {

	err := repo.Conn.Create(&input).Error

	if err != nil {
		panic("Failed to create Product: " + err.Error())
	}

	return input, err
}
