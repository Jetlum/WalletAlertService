package services

import (
	"fmt"

	"github.com/Jetlum/WalletAlertService/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailNotification struct {
	client    *sendgrid.Client
	fromEmail string
}

func NewEmailNotification(apiKey string) *EmailNotification {
	return &EmailNotification{
		client:    sendgrid.NewSendClient(apiKey),
		fromEmail: "alerts@walletalert.service",
	}
}

func (en *EmailNotification) Send(event *models.Event, userPref *models.UserPreference) error {
	if event == nil || userPref == nil {
		return fmt.Errorf("event and user preference cannot be nil")
	}

	from := mail.NewEmail("Wallet Alert Service", en.fromEmail)
	to := mail.NewEmail("", userPref.UserID)
	subject := fmt.Sprintf("Alert: %s Event Detected", event.EventType)
	content := formatEventMessage(event)

	message := mail.NewSingleEmail(from, subject, to, content, content)
	response, err := en.client.Send(message)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("email API error: status %d", response.StatusCode)
	}

	return nil
}
