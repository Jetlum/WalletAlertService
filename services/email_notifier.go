package services

import "github.com/Jetlum/WalletAlertService/models"

type EmailNotifier interface {
	Send(event *models.Event, userPref *models.UserPreference) error
}

type EmailNotification struct {
	APIKey string
}

var _ EmailNotifier = (*EmailNotification)(nil)

func (en *EmailNotification) Send(event *models.Event, userPref *models.UserPreference) error {
	return nil
}
