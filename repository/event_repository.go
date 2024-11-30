package repository

import (
	"fmt"

	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) Create(event *models.Event) error {
	if event == nil {
		return fmt.Errorf("event cannot be nil")
	}
	return er.db.Create(event).Error
}

func (er *EventRepository) GetUnnotifiedEvents() ([]models.Event, error) {
	var events []models.Event
	err := er.db.Where("notified = ?", false).Find(&events).Error
	return events, err
}

func (er *EventRepository) MarkAsNotified(eventID uint) error {
	return er.db.Model(&models.Event{}).Where("id = ?", eventID).Update("notified", true).Error
}
