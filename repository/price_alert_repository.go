package repository

import (
	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type PriceAlertRepository struct {
	db *gorm.DB
}

func NewPriceAlertRepository(db *gorm.DB) *PriceAlertRepository {
	return &PriceAlertRepository{db: db}
}

func (r *PriceAlertRepository) Create(alert *models.PriceAlert) error {
	return r.db.Create(alert).Error
}

func (r *PriceAlertRepository) GetActiveAlerts() ([]models.PriceAlert, error) {
	var alerts []models.PriceAlert
	err := r.db.Where("is_active = ?", true).Find(&alerts).Error
	return alerts, err
}

func (r *PriceAlertRepository) DeactivateAlert(alertID uint) error {
	return r.db.Model(&models.PriceAlert{}).Where("id = ?", alertID).
		Update("is_active", false).Error
}
