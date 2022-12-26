package repository

import (
	"github.com/kleberalves/problemCompanyApp/backend/purchase"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"gorm.io/gorm"
)

type purchaseRepository struct {
	Conn *gorm.DB
}

func NewPurchaseRepository(Conn *gorm.DB) purchase.Repository {
	return &purchaseRepository{Conn}
}

func (repo *purchaseRepository) FindAll() (res []schema.Purchase, err error) {

	var purchases []schema.Purchase
	errExec := repo.Conn.Model(&schema.Purchase{}).Find(&purchases).Error

	if errExec != nil {
		panic("Failed to retrieve all Purchases: " + err.Error())
	}

	return purchases, errExec
}

func (repo *purchaseRepository) Create(input schema.Purchase) (schema.Purchase, error) {

	err := repo.Conn.Create(&input).Error

	return input, err
}
