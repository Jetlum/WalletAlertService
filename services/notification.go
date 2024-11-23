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
	return fmt.Sprintf(
		"Transaction detected:\nFrom: %s\nTo: %s\nValue: %s\nType: %s",
		event.FromAddress,
		event.ToAddress,
		event.Value,
		event.EventType,
	)
}
