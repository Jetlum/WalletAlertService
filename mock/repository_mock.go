package mock

import "github.com/Jetlum/WalletAlertService/models"

// Mock Event Repository
type MockEventRepository struct {
	CreateFunc func(event *models.Event) error
}

func NewMockEventRepository() *MockEventRepository {
	return &MockEventRepository{
		CreateFunc: func(event *models.Event) error {
			return nil
		},
	}
}

func (m *MockEventRepository) Create(event *models.Event) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(event)
	}
	return nil
}

// Mock User Preference Repository
type MockUserPreferenceRepository struct {
	GetMatchingPreferencesFunc func(event *models.Event) ([]models.UserPreference, error)
}

func NewMockUserPreferenceRepository() *MockUserPreferenceRepository {
	return &MockUserPreferenceRepository{
		GetMatchingPreferencesFunc: func(event *models.Event) ([]models.UserPreference, error) {
			return []models.UserPreference{}, nil
		},
	}
}

func (m *MockUserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	if m.GetMatchingPreferencesFunc != nil {
		return m.GetMatchingPreferencesFunc(event)
	}
	return []models.UserPreference{}, nil
}
