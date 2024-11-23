package mock

import (
	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type MockDB struct {
	Events          []models.Event
	UserPreferences []models.UserPreference
}

func NewMockDB() *MockDB {
	return &MockDB{
		Events:          []models.Event{},
		UserPreferences: []models.UserPreference{},
	}
}

func (db *MockDB) Create(value interface{}) *gorm.DB {
	switch v := value.(type) {
	case *models.Event:
		db.Events = append(db.Events, *v)
	case *models.UserPreference:
		db.UserPreferences = append(db.UserPreferences, *v)
	}
	return &gorm.DB{}
}

func (db *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	switch out := out.(type) {
	case *[]models.Event:
		*out = db.Events
	case *[]models.UserPreference:
		*out = db.UserPreferences
	}
	return &gorm.DB{}
}
