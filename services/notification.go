package services

import (
	"fmt"
	"models"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type NotificationService interface {
	Send(event *models.Event, userPreference *models.UserPreference) error
}

type EmailNotification struct {
	apiKey string
}

func NewEmailNotification(apiKey string) *EmailNotification {
	return &EmailNotification{apiKey: apiKey}
}

func (n *EmailNotification) Send(event *models.Event, userPref *models.UserPreference) error {
	from := mail.NewEmail("Wallet Alert", "alerts@yourdomain.com")
	subject := "New Wallet Activity Detected"
	to := mail.NewEmail("User", userPref.UserID)
	content := mail.NewContent("text/plain", formatEventMessage(event))

	message := mail.NewV3MailInit(from, subject, to, content)
	client := sendgrid.NewSendClient(n.apiKey)

	_, err := client.Send(message)
	return err
}

func formatEventMessage(event *models.Event) string {
	return fmt.Sprintf(
		"Transaction detected:\nFrom: %s\nTo: %s\nValue: %s\nType: %s",
		event.FromAddress,
		event.ToAddress,
		event.Value.String(),
		event.EventType,
	)
}
