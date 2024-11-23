package models

import (
	"gorm.io/gorm"
)

type UserPreference struct {
	gorm.Model
	UserID            string `gorm:"uniqueIndex"`
	WalletAddress     string `gorm:"index"`
	MinEtherValue     string `gorm:"type:numeric"`
	TrackNFTs         bool
	EmailNotification bool
	PushNotification  bool
}
