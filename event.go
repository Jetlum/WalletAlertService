package event

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	TxHash      string `gorm:"uniqueIndex"`
	FromAddress string `gorm:"index"`
	ToAddress   string `gorm:"index"`
	Value       string `gorm:"type:numeric"`
	EventType   string `gorm:"index"`
	Notified    bool
}
