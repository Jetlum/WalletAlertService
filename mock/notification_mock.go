package mock

import "github.com/Jetlum/WalletAlertService/models"

type MockEmailNotification struct {
	SendFunc func(event *models.Event, userPref *models.UserPreference) error
}

func (m *MockEmailNotification) Send(event *models.Event, userPref *models.UserPreference) error {
	if m.SendFunc != nil {
		return m.SendFunc(event, userPref)
	}
	return nil
}
