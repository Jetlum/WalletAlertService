package mock

import (
	"github.com/Jetlum/WalletAlertService/models"
	"github.com/Jetlum/WalletAlertService/services"
)

type MockEmailNotification struct {
	services.EmailNotifier
	SendFunc func(event *models.Event, userPref *models.UserPreference) error
}

func (m *MockEmailNotification) Send(event *models.Event, userPref *models.UserPreference) error {
	if m.SendFunc != nil {
		return m.SendFunc(event, userPref)
	}
	return nil
}
