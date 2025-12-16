package database

import (
	"fmt"
	"log"
	"test-webdev-suiten-25/config"

	// Driver GORM untuk PostgreSQL
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	dsn := cfg.GetDBConnectionString()

	dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database with GORM: %w", err)
	}

	DB = dbGorm

	log.Println("Database connected successfully using GORM")

	return nil
}
