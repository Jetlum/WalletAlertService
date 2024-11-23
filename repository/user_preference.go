package repository

import (
	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type UserPreferenceRepository struct {
	db *gorm.DB
}

func NewUserPreferenceRepository(db *gorm.DB) *UserPreferenceRepository {
	return &UserPreferenceRepository{db: db}
}

func (r *UserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	var preferences []models.UserPreference
	err := r.db.Where("wallet_address = ?", event.ToAddress).Find(&preferences).Error
	return preferences, err
}
