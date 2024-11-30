package main

import (
	"os"
	"testing"

	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Set test environment
	os.Setenv("GO_ENV", "test")
	// Run tests
	os.Exit(m.Run())
}

func TestNotifyUsers(t *testing.T) {
	event := &models.Event{
		TxHash:      "0xabc123",
		FromAddress: "0xfrom",
		ToAddress:   "0xto",
		Value:       "1000000000000000000",
		EventType:   "LARGE_TRANSFER",
	}

	emailSent := false
	mockUserPrefRepo := &mock.MockUserPreferenceRepository{
		GetMatchingPreferencesFunc: func(event *models.Event) ([]models.UserPreference, error) {
			return []models.UserPreference{{
				UserID:            "user@example.com",
				WalletAddress:     event.ToAddress,
				EmailNotification: true,
			}}, nil
		},
	}

	mockEmailNotification := &mock.MockEmailNotification{
		SendFunc: func(event *models.Event, userPref *models.UserPreference) error {
			emailSent = true
			return nil
		},
	}

	notifyUsers(event, mockUserPrefRepo, mockEmailNotification)
	assert.True(t, emailSent, "Email notification should have been sent")
}
