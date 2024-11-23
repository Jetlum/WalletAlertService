package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/Jetlum/WalletAlertService/mock"
	"github.com/Jetlum/WalletAlertService/models"
)

func TestIsLargeTransfer(t *testing.T) {
	tests := []struct {
		name     string
		value    *big.Int
		expected bool
	}{
		{
			name:     "Large transfer",
			value:    big.NewInt(2000000000000000000), // 2 ETH
			expected: true,
		},
		{
			name:     "Small transfer",
			value:    big.NewInt(500000000000000000), // 0.5 ETH
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := types.NewTransaction(
				0,
				common.Address{},
				tt.value,
				21000,
				big.NewInt(1),
				nil,
			)
			result := isLargeTransfer(tx)
			assert.Equal(t, tt.expected, result)
		})
	}
}

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

	var emailSent bool
	mockEmailNotification := &mock.MockEmailNotification{
		SendFunc: func(e *models.Event, up *models.UserPreference) error {
			emailSent = true
			return nil
		},
	}

	notifyUsers(event, mockUserPrefRepo, mockEmailNotification)
	assert.True(t, emailSent, "Email notification should have been sent")
}
