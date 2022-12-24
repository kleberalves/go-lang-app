package models

import (
	"github.com/kleberalves/problemCompanyApp/backend/datasources"
)

func AutoMigrations() {

	db := datasources.Connect()

	db.AutoMigrate(&User{}, &Product{}, &Purchase{})

	// err := datasources.DB.AutoMigrate(&User{})

	// if err != nil {
	// 	return
	// }
}
