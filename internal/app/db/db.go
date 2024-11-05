package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect(url string) *gorm.DB {
	var err error
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
	return db
}
