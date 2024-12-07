package database

import (
	"testing"

	"github.com/Jetlum/WalletAlertService/config"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	cfg, err := config.LoadConfig()
	assert.NoError(t, err)

	err = InitDB(cfg.DatabaseURL)
	assert.NoError(t, err)
	defer CloseDB()

	// Test database connection
	sqlDB, err := DB.DB()
	assert.NoError(t, err)
	assert.NoError(t, sqlDB.Ping())
}
