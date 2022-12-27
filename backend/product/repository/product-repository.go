package repository

import (
	"github.com/kleberalves/problemCompanyApp/backend/product"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"gorm.io/gorm"
)

type repository struct {
	Conn *gorm.DB
}

func NewProductRepository(Conn *gorm.DB) product.Repository {
	return &repository{Conn}
}

func (repo *repository) FindAll() (res []schema.Product, err error) {

	var products []schema.Product
	errExec := repo.Conn.Model(&schema.Product{}).Find(&products).Error

	if errExec != nil {
		panic("Failed to retrieve all Products: " + err.Error())
	}

	return products, errExec
}

func (repo *repository) Create(input schema.Product) (schema.Product, error) {

	err := repo.Conn.Create(&input).Error

	if err != nil {
		panic("Failed to create Product: " + err.Error())
	}

	return input, err
}

func (repo *repository) Update(item schema.Product) error {
	err := repo.Conn.Model(&schema.Product{}).
		Where("id = ? ", item.ID).
		Updates(&item).Error
	return err
}

func (repo *repository) UpdateDeletedAt(item schema.Product) error {
	err := repo.Conn.Model(&schema.Product{}).
		Where("id = ? ", item.ID).
		Update("deleted_at", nil).Error
	return err
}

func (repo *repository) Delete(itemIds []int) error {
	//hard deletes
	err := repo.Conn.Unscoped().Delete(&schema.Product{}, itemIds).Error
	return err
}
