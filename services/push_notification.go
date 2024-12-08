// services/push_notification.go
package services

import "github.com/Jetlum/WalletAlertService/models"

type PushNotifier interface {
	Send(event *models.Event, userPref *models.UserPreference) error
}

type PushNotification struct {
	// Firebase or other push service configuration
}

func NewPushNotification(config Config) *PushNotification {
	return &PushNotification{
		// Initialize push service
	}
}

func (pn *PushNotification) Send(event *models.Event, userPref *models.UserPreference) error {
	// Implement push notification sending
	return nil
}
