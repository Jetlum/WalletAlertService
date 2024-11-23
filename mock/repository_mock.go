package mock

import "github.com/Jetlum/WalletAlertService/models"

type MockUserPreferenceRepository struct {
	GetMatchingPreferencesFunc func(event *models.Event) ([]models.UserPreference, error)
}

func (m *MockUserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	if m.GetMatchingPreferencesFunc != nil {
		return m.GetMatchingPreferencesFunc(event)
	}
	return nil, nil
}
