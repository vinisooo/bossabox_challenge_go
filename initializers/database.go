package initializers

import (
	"gorm.io/driver/postgres"
	"log"
	"os"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connet to database")
	}
}