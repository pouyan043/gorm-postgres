package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}
