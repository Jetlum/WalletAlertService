// repository/interfaces.go
package repository

import "github.com/Jetlum/WalletAlertService/models"

type UserPreferenceRepository interface {
	GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error)
}

type EventRepository interface {
	Create(event *models.Event) error
}
