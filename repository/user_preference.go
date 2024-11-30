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

func (r *UserPreferenceRepository) GetMatchingPreferences(event *models.Event) ([]models.UserPreference, error) {
	var preferences []models.UserPreference

	eventValue, ok := new(big.Int).SetString(event.Value, 10)
	if !ok {
		return nil, fmt.Errorf("invalid event value: %s", event.Value)
	}

	// Base query
	query := r.db.Where("wallet_address = ?", event.ToAddress)

	// Add filters based on event type
	switch event.EventType {
	case "LARGE_TRANSFER":
		query = query.Where("min_ether_value <= ?", eventValue.String())
	case "NFT_TRANSFER":
		query = query.Where("track_nfts = ?", true)
	}

	if err := query.Find(&preferences).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch preferences: %w", err)
	}

	return preferences, nil
}
