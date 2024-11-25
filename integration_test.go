package main

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Setup test context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create test event
	event := &models.Event{
		TxHash:      "0xtest",
		FromAddress: "0xfrom",
		ToAddress:   "0xto",
		Value:       "1000000000000000000", // 1 ETH
		EventType:   "LARGE_TRANSFER",
	}

	// Setup mocks
	mockEventRepo := &mock.MockEventRepository{
		CreateFunc: func(e *models.Event) error {
			assert.Equal(t, event.TxHash, e.TxHash)
			return nil
		},
	}

	userPref := models.UserPreference{
		UserID:            "test@example.com",
		WalletAddress:     "0xto",
		MinEtherValue:     "500000000000000000", // 0.5 ETH
		EmailNotification: true,
	}

	mockUserPrefRepo := &mock.MockUserPreferenceRepository{
		GetMatchingPreferencesFunc: func(e *models.Event) ([]models.UserPreference, error) {
			assert.Equal(t, event.ToAddress, e.ToAddress)
			return []models.UserPreference{userPref}, nil
		},
	}

	// Setup mock email notification
	emailSent := false
	mockEmailNotification := &mock.MockEmailNotification{
		SendFunc: func(e *models.Event, up *models.UserPreference) error {
			emailSent = true
			return nil
		},
	}

	// Create mock block and header
	mockHeader := &types.Header{
		Number: big.NewInt(1),
	}

	// Test full workflow
	done := make(chan bool)
	go func() {
		processBlock(
			nil,
			mockHeader,
			mock.NewMockNFTDetector(),
			mockEmailNotification,
			mockEventRepo,
			mockUserPrefRepo,
		)
		done <- true
	}()

	// Wait for completion or timeout
	select {
	case <-ctx.Done():
		t.Fatal("Test timed out")
	case <-done:
		assert.True(t, emailSent, "Email notification should have been sent")
	}
}
