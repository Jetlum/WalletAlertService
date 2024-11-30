package services

import (
	"fmt"

	"github.com/Jetlum/WalletAlertService/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailNotifier interface {
	Send(event *models.Event, userPref *models.UserPreference) error
}

type EmailNotification struct {
	APIKey string
	client *sendgrid.Client
}

func (en *EmailNotification) Send(event *models.Event, userPref *models.UserPreference) error {
	if userPref.UserID == "" {
		return fmt.Errorf("invalid user email")
	}

	from := mail.NewEmail("Wallet Alert Service", "alerts@walletalert.service")
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
