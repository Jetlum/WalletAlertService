package repository

import "github.com/Jetlum/WalletAlertService/models"

type EventRepositoryInterface interface {
	Create(event *models.Event) error
}

type UserPreferenceRepositoryInterface interface {
	GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error)
}
