package schema

import (
	"gorm.io/gorm"
)

// type tableInfo struct {
// 	table_name  string
// 	column_name string
// 	data_type   string
// }

func AutoMigrations(db *gorm.DB) {

	db.AutoMigrate(&User{}, &Product{}, &Purchase{}, &Profile{})

}
