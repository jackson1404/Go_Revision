package db_configs

import (
	"fmt"

	"jackson.com/libraryapisystem/Configurations/env_configs"
	"jackson.com/libraryapisystem/Configurations/logger_configs"
	"jackson.com/libraryapisystem/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	cfg := env_configs.AppConfig

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger_configs.Errorw("❌ Failed to connect to database", "error", err)
		return err
	}

	DB = db
	logger_configs.Info("Database connected successfully!")

	if err := AutoMigrate(); err != nil {
		logger_configs.Errorw("Failed to run migrations", "error", err)
		return err
	}

	logger_configs.Info("Database migrated successfully!")
	return nil
}

// AutoMigrate all models
func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Author{},
		&models.Book{},
		&models.Loan{},
		&models.Reservation{},
	)
}
