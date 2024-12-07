package main

import (
	"os"
	"testing"

	"github.com/Jetlum/WalletAlertService/config"
	"github.com/Jetlum/WalletAlertService/database"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/Jetlum/WalletAlertService/repository"
	"github.com/Jetlum/WalletAlertService/services"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	os.Exit(m.Run())
}

func TestNotifyUsers(t *testing.T) {
	// Initialize with real components
	cfg, err := config.LoadConfig()
	assert.NoError(t, err)

	err = database.InitDB(cfg.DatabaseURL)
	assert.NoError(t, err)
	defer database.CloseDB()

	// Create real repositories
	eventRepo := repository.NewEventRepository(database.DB)
	userPrefRepo := repository.NewUserPreferenceRepository(database.DB)
	emailNotification := services.NewEmailNotification(cfg.SendGridAPIKey)

	// Test with real event
	event := &models.Event{
		TxHash:      "0xabc123",
		FromAddress: "0xfrom",
		ToAddress:   "0xto",
		Value:       "1000000000000000000",
		EventType:   "LARGE_TRANSFER",
	}

	// Save event to database
	err = eventRepo.Create(event)
	assert.NoError(t, err)

	// Test notification
	err = notifyUsers(event, userPrefRepo, emailNotification)
	assert.NoError(t, err)
}
