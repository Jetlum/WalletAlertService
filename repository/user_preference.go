package repository

import (
	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type UserPreferenceRepository struct {
	db *gorm.DB
}

var _ UserPreferenceRepositoryInterface = (*UserPreferenceRepository)(nil)

func NewUserPreferenceRepository(db *gorm.DB) *UserPreferenceRepository {
	return &UserPreferenceRepository{db: db}
}

func (upr *UserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	return nil, nil
}
