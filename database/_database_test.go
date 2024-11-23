package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	SetupMockDB()
	defer ResetMockDB()

	t.Run("Mock Mode", func(t *testing.T) {
		db, err := InitDB("any-dsn")
		assert.Nil(t, err)
		assert.Nil(t, db)
	})

	t.Run("Empty DSN", func(t *testing.T) {
		IsMockMode = false
		db, err := InitDB("")
		assert.Error(t, err)
		assert.Nil(t, db)
	})
}
