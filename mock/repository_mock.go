package mock

import (
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/Jetlum/WalletAlertService/repository"
)

type MockUserPreferenceRepository struct {
	repository.UserPreferenceRepositoryInterface
	GetMatchingPreferencesFunc func(event *models.Event) ([]models.UserPreference, error)
}

func (m *MockUserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	if m.GetMatchingPreferencesFunc != nil {
		return m.GetMatchingPreferencesFunc(event)
	}
	return nil, nil
}

type MockEventRepository struct {
	repository.EventRepositoryInterface
	CreateFunc func(event *models.Event) error
}

func (m *MockEventRepository) Create(event *models.Event) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(event)
	}
	return nil
}
