package datasources

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostGresDataSource() DataSource {
	return &dataSource{}
}

type dataSource struct {
}

func (ds *dataSource) Connect() *gorm.DB {

	dbStr := os.ExpandEnv("host=$POSTGRES_HOSTNAME user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB port=$POSTGRES_PORT sslmode=disable TimeZone=America/Sao_Paulo")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	return db

}
