package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Exported DB variable
var DB *gorm.DB

func InitDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
