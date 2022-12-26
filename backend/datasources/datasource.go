package datasources

import "gorm.io/gorm"

type DataSource interface {
	Connect() *gorm.DB
}
