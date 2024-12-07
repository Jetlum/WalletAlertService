package models

import (
	"gorm.io/gorm"
)

// Add new model for price alerts
type PriceAlert struct {
	gorm.Model
	UserID            string `gorm:"index"`
	CryptocurrencyID  string `gorm:"index"`        // e.g., "BTC", "ETH"
	ThresholdPrice    string `gorm:"type:numeric"` // Store as string for precision
	IsUpperBound      bool   // true for price above, false for below
	IsActive          bool   `gorm:"default:true"`
	EmailNotification bool   `gorm:"default:true"`
}

// Update UserPreference to include price alerts
type UserPreference struct {
	gorm.Model
	UserID            string `gorm:"uniqueIndex"`
	WalletAddress     string `gorm:"index"`
	MinEtherValue     string `gorm:"type:numeric"`
	TrackNFTs         bool
	EmailNotification bool
	PushNotification  bool
	PriceAlerts       []PriceAlert `gorm:"foreignKey:UserID"`
}
