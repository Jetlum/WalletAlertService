package main

import (
	"testing"

	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/stretchr/testify/assert"
)

func TestNotifyUsers(t *testing.T) {
	// Create a sample event to pass to notifyUsers
	event := &models.Event{
		TxHash:      "0xabc123",
		FromAddress: "0xfrom",
		ToAddress:   "0xto",
		Value:       "1000000000000000000", // 1 ETH in Wei
		EventType:   "LARGE_TRANSFER",
	}

	// Initialize the mock repositories and services
	mockUserPrefRepo := &mock.MockUserPreferenceRepository{
		GetMatchingPreferencesFunc: func(event *models.Event) ([]models.UserPreference, error) {
			return []models.UserPreference{
				{
					UserID:            "user@example.com",
					WalletAddress:     event.ToAddress,
					EmailNotification: true,
				},
			}, nil
		},
	}

	emailSent := false

	mockEmailNotification := &mock.MockEmailNotification{
		SendFunc: func(event *models.Event, userPref *models.UserPreference) error {
			emailSent = true
			return nil
		},
	}

	// Call the function under test
	notifyUsers(event, mockUserPrefRepo, mockEmailNotification)

	// Assert that email was sent
	assert.True(t, emailSent, "Email notification should have been sent")
}
