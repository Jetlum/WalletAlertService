package repository

import (
	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

var _ EventRepositoryInterface = (*EventRepository)(nil)

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) Create(event *models.Event) error {
	return er.db.Create(event).Error
}

func (er *EventRepository) GetUnnotifiedEvents() ([]models.Event, error) {
	return nil, nil
}

func (er *EventRepository) MarkAsNotified(eventID uint) error {
	return nil
}
