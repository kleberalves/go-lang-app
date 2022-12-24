package datasources

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dbStr := os.ExpandEnv("host=$POSTGRES_HOSTNAME user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB port=$POSTGRES_PORT sslmode=disable TimeZone=America/Sao_Paulo")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbStr, // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,  // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	return db

}
