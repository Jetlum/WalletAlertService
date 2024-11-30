package database

import (
	"fmt"

	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations
	err = DB.AutoMigrate(&models.Event{}, &models.UserPreference{})
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
