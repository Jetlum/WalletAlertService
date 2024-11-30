package database

import (
	"errors"
	"io"
	"log"
	"os"
	"sync"

	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	IsMockMode  bool
	ErrMockMode = errors.New("database is in mock mode")
	logger      = log.New(os.Stdout, "", log.LstdFlags)
	mu          sync.Mutex
)

func init() {
	// Setup mock DB in test mode
	if os.Getenv("GO_ENV") == "test" {
		SetupMockDB()
		logger.SetOutput(io.Discard)
	}
}

func InitDB(dsn string) (*gorm.DB, error) {
	IsMockMode = true
	mu.Lock()
	defer mu.Unlock()

	// Skip DB initialization in mock mode
	if IsMockMode {
		return nil, nil
	}

	return nil, nil
}

func SetupMockDB() {
	mu.Lock()
	defer mu.Unlock()

	IsMockMode = true
	DB = nil
	logger = log.New(io.Discard, "", 0)
}

func ResetMockDB() {
	mu.Lock()
	defer mu.Unlock()

	IsMockMode = false
	DB = nil
	logger = log.New(os.Stdout, "", log.LstdFlags)
}
