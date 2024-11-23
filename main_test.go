// main_test.go
package main

import (
	"testing"

	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/stretchr/testify/assert"
)

func TestNotifyUsers(t *testing.T) {
	event := &models.Event{
		TxHash:      "0x123",
		FromAddress: "0x456",
		ToAddress:   "0x789",
		Value:       "1000000000000000000",
		EventType:   "LARGE_TRANSFER",
	}

	mockUserPrefRepo := &mock.MockUserPreferenceRepository{
		GetMatchingPreferencesFunc: func(e *models.Event) ([]models.UserPreference, error) {
			return []models.UserPreference{
				{
					UserID:            "test@example.com",
					WalletAddress:     "0x789",
					EmailNotification: true,
				},
			}, nil
		},
	}

	emailSent := false
	mockEmailNotification := &mock.MockEmailNotification{
		SendFunc: func(e *models.Event, up *models.UserPreference) error {
			emailSent = true
			return nil
		},
	}

	notifyUsers(event, mockUserPrefRepo, mockEmailNotification)
	assert.True(t, emailSent, "Email notification should have been sent")
}
