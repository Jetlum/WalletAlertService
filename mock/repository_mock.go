package mock

import (
	"github.com/Jetlum/WalletAlertService/models"
)

type MockEventRepository struct {
	CreateFunc func(event *models.Event) error
}

type MockUserPreferenceRepository struct {
	GetMatchingPreferencesFunc func(event *models.Event) ([]models.UserPreference, error)
}

func (m *MockEventRepository) Create(event *models.Event) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(event)
	}
	return nil
}

func (m *MockUserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	if m.GetMatchingPreferencesFunc != nil {
		return m.GetMatchingPreferencesFunc(event)
	}
	return nil, nil
}
