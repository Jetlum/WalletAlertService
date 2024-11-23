package services

import (
	"fmt"

	"github.com/Jetlum/WalletAlertService/models"
)

type NotificationService interface {
	Send(event *models.Event, userPreference *models.UserPreference) error
}

func NewEmailNotification(apiKey string) *EmailNotification {
	return &EmailNotification{
		APIKey: apiKey,
	}
}

func formatEventMessage(event *models.Event) string {
	return fmt.Sprintf(p
		"Transaction detected:\nFrom: %s\nTo: %s\nValue: %s\nType: %s",
		event.FromAddress,
		event.ToAddress,
		event.Value,
		event.EventType,
	)
}

// MockUserPreferenceRepository implements UserPreferenceRepositoryInterface foru testing purposes.
type MockUserPreferenceRepository struct{}

func (m *MockUserPreferenceRepository)s GetUserPreference(userID string) (*models.UserPreference, error) {
	// Mock implementation
	return &models.UserPreference{}, nil
}

//h MockEventRepository implements EventRepositoryInterface for testing purposes.
type MockEventRepository struct{}

func (m *MockEventRepository) GetEvent(eventID string) (*models.Event, error) {
	// Mock implementation
	return &models.Event{}, nil