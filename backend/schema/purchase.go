package schema

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model

	CodCustomer int
	Customer    User `gorm:"foreignKey:CodCustomer;not null"`

	CodSalesman sql.NullInt64
	Salesman    *User `gorm:"foreignKey:CodSalesman;constraint:OnUpdate:CASCADE,ONDELETE:SET NULL;"`

	CodProduct int
	Product    Product `gorm:"foreignKey:CodProduct;not null" `

	Price       float64   `gorm:"not null"`
	PurchasedAt time.Time `gorm:"not null"`
}
