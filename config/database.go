package config

import (
	"fmt"
	"log"

	"hello/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := fmt.Sprintf(
		"host=localhost user=postgres password=Venkey83$ dbname=project port=5432 sslmode=disable",
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Database connection established")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate database: %w", err)
	}
	fmt.Println("Database auto-migration completed")
	return nil
}
