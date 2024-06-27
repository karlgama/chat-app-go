package postgreSQL

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func getConnectionString() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || name == "" {
		log.Fatalf("One or more environment variables are missing")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port)
}

func SetupDatabase() {
	connString := getConnectionString()
	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database", err)
	}
}
