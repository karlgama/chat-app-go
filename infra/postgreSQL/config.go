package postgreSQL

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
	DB  *gorm.DB
	err error
)

func SetupDatabase() {
	connString := os.Getenv("POSTGRE_CONNECTION_STRING")
	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database", err)
	}
}
