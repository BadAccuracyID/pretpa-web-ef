package database

import (
	"fmt"
	"github.com/badaccuracyid/pretpa-web-ef/internal/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	database *gorm.DB
	once     sync.Once
)

func GetDatabase() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		database, err = connect()
	})

	return database, err
}

func MigrateTables() (bool, error) {
	db, err := GetDatabase()
	if err != nil {
		return false, err
	}

	err = db.AutoMigrate(&model.User{})
	err = db.AutoMigrate(&model.Conversation{})
	err = db.AutoMigrate(&model.Message{})
	if err != nil {
		return false, fmt.Errorf("failed to migrate tables: %w", err)
	}

	return true, nil
}

func connect() (*gorm.DB, error) {
	envDsn := os.Getenv("DATABASE_DSN")
	if envDsn == "" {
		return nil, fmt.Errorf("DATABASE_DSN is not set")
	}

	dsn := envDsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db, nil
}
