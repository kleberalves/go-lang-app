package schema

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model

	CodCustomer int  `json:"codCustomer"`
	Customer    User `gorm:"foreignKey:CodCustomer"`

	CodSalesman int  `json:"codSalesman"`
	Salesman    User `gorm:"foreignKey:CodSalesman"`

	CodProduct int     `json:"codProduct"`
	Product    Product `gorm:"foreignKey:CodProduct"`

	Price       float64   `json:"price"`
	PurchasedAt time.Time `json:"date"`
}
