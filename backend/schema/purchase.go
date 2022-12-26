package schema

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model

	CodCustomer int
	Customer    User `gorm:"foreignKey:CodCustomer;not null"`

	CodSalesman int
	Salesman    User `gorm:"foreignKey:CodSalesman;not null"`

	CodProduct int
	Product    Product `gorm:"foreignKey:CodProduct;not null" `

	Price       float64   `gorm:"not null"`
	PurchasedAt time.Time `gorm:"not null"`
}
