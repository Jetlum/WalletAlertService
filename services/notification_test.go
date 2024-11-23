package services

import (
	"testing"

	"github.com/Jetlum/WalletAlertService/models"
	"github.com/stretchr/testify/assert"
)

func TestFormatEventMessage(t *testing.T) {
	event := &models.Event{
		FromAddress: "0x123",
		ToAddress:   "0x456",
		Value:       "1000000000000000000",
		EventType:   "LARGE_TRANSFER",
	}

	message := formatEventMessage(event)
	expected := "Transaction detected:\nFrom: 0x123\nTo: 0x456\nValue: 1000000000000000000\nType: LARGE_TRANSFER"

	assert.Equal(t, expected, message)
}
