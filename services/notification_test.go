package services

import (
	"testing"

	"github.com/Jetlum/WalletAlertService/config"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/stretchr/testify/assert"
)

func TestFormatEventMessage(t *testing.T) {
	// Get config
	cfg, err := config.LoadConfig()
	assert.NoError(t, err)

	// Initialize email notification service
	emailNotification := NewEmailNotification(cfg.SendGridAPIKey)

	// Create test event
	event := &models.Event{
		FromAddress: "0x123",
		ToAddress:   "0x456",
		Value:       "1000000000000000000",
		EventType:   "LARGE_TRANSFER",
	}

	// Create test user preference
	userPref := &models.UserPreference{
		UserID:            "test@example.com",
		EmailNotification: true,
	}

	// Test sending notification
	err = emailNotification.Send(event, userPref)
	assert.NoError(t, err)
}
