package repository

import (
	"github.com/kleberalves/problemCompanyApp/backend/purchase"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"gorm.io/gorm"
)

type repository struct {
	Conn *gorm.DB
}

func NewPurchaseRepository(Conn *gorm.DB) purchase.Repository {
	return &repository{Conn}
}

func (repo *repository) FindAll() (res []schema.Purchase, err error) {

	var purchases []schema.Purchase
	errExec := repo.Conn.Model(&schema.Purchase{}).
		Preload("Customer", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Salesman", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Product", &schema.Product{}).
		Order("purchased_at desc").
		Find(&purchases).Error

	return purchases, errExec
}

func (repo *repository) GetByUser(userId int) ([]schema.Purchase, error) {
	var purchases []schema.Purchase
	errExec := repo.Conn.Model(&schema.Purchase{}).
		Preload("Salesman", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Product", &schema.Product{}).
		Where("cod_customer = ? ", userId).
		Order("purchased_at desc").
		Find(&purchases).Error
	return purchases, errExec
}

func (repo *repository) Create(input schema.Purchase) (schema.Purchase, error) {

	err := repo.Conn.
		Preload("Customer", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Salesman", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Product", &schema.Product{}).
		Create(&input).
		Error

	return input, err
}

func (repo *repository) Delete(purchaseIds []int) error {
	// Unscoped() deletes permanently
	err := repo.Conn.Unscoped().
		Delete(&schema.Purchase{},
			repo.Conn.Where(purchaseIds)).
		Error
	return err
}
