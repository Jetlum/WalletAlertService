package repository

import (
	"github.com/Jetlum/WalletAlertService/database"
	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		db: database.DB,
	}
}

func (r *EventRepository) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *EventRepository) GetUnnotifiedEvents() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Where("notified = ?", false).Find(&events).Error
	return events, err
}

func (r *EventRepository) MarkAsNotified(eventID uint) error {
	return r.db.Model(&models.Event{}).Where("id = ?", eventID).Update("notified", true).Error
}
