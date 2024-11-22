package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"interface.social/models"
)

var DB *gorm.DB

func InitDB(dbURL string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}

	// Run migrations
	err = DB.AutoMigrate(&models.UserPreference{}, &models.Event{})
	if err != nil {
		return err
	}

	return nil
}
