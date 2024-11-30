package repository

import (
	"fmt"
	"math/big"

	"github.com/Jetlum/WalletAlertService/models"
	"gorm.io/gorm"
)

type UserPreferenceRepository struct {
	db *gorm.DB
}

func NewUserPreferenceRepository(db *gorm.DB) *UserPreferenceRepository {
	return &UserPreferenceRepository{db: db}
}

func (upr *UserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	if event == nil {
		return nil, fmt.Errorf("event cannot be nil")
	}

	var preferences []models.UserPreference
	query := upr.db.Where("wallet_address = ?", event.ToAddress)

	if event.EventType == "LARGE_TRANSFER" {
		eventValue, ok := new(big.Int).SetString(event.Value, 10)
		if !ok {
			return nil, fmt.Errorf("invalid event value: %s", event.Value)
		}
		query = query.Where("min_ether_value <= ?", eventValue.String())
	} else if event.EventType == "NFT_TRANSFER" {
		query = query.Where("track_nfts = ?", true)
	}

	err := query.Find(&preferences).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch preferences: %w", err)
	}

	return preferences, nil
}
