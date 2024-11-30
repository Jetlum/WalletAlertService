package mock

import (
	"github.com/Jetlum/WalletAlertService/models"
)

type MockDB struct {
	Events          []models.Event
	UserPreferences []models.UserPreference
}

var TestDB = &MockDB{
	Events:          make([]models.Event, 0),
	UserPreferences: make([]models.UserPreference, 0),
}

func ResetTestDB() {
	TestDB.Events = make([]models.Event, 0)
	TestDB.UserPreferences = make([]models.UserPreference, 0)
}
