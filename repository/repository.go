package repository

import "github.com/Jetlum/WalletAlertService/models"

type UserPreferenceRepositoryInterface interface {
	GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error)
}

type EventRepositoryInterface interface {
	Create(event *models.Event) error
	GetUnnotifiedEvents() ([]models.Event, error)
	MarkAsNotified(eventID uint) error
}
