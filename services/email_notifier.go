package services

import (
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailNotifier interface {
	Send(event *models.Event, userPref *models.UserPreference) error
}

type EmailNotification struct {
	APIKey string
}

var _ EmailNotifier = (*EmailNotification)(nil)

func (en *EmailNotification) Send(event *models.Event, userPref *models.UserPreference) error {
	from := mail.NewEmail("Wallet Alert Service", "no-reply@example.com")
	to := mail.NewEmail("", userPref.UserID)
	subject := "Wallet Activity Alert"
	content := formatEventMessage(event)
	message := mail.NewSingleEmail(from, subject, to, content, content)

	client := sendgrid.NewSendClient(en.APIKey)
	_, err := client.Send(message)
	return err
}
