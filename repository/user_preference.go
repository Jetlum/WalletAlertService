package repository

import (
	"github.com/Jetlum/WalletAlertService/database"
	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type UserPreferenceRepository struct {
	db *gorm.DB
}

func NewUserPreferenceRepository() *UserPreferenceRepository {
	return &UserPreferenceRepository{
		db: database.DB,
	}
}

func (r *UserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	var preferences []models.UserPreference
	err := r.db.Where("wallet_address = ? AND min_ether_value <= ?", event.ToAddress, event.Value).Find(&preferences).Error
	return preferences, err
}
