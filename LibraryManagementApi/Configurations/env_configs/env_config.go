package env_configs

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"jackson.com/libraryapisystem/Configurations/logger_configs"
)

type Config struct {
	DB_HOST     string `env:"DB_HOST,required"`
	DB_USER     string `env:"DB_USER,required"`
	DB_PASSWORD string `env:"DB_PASSWORD,required"`
	DB_NAME     string `env:"DB_NAME,required"`
	DB_PORT     int    `env:"DB_PORT" envDefault:"5432"`

	APP_PORT   int    `env:"APP_PORT" envDefault:"8080"`
	API_PREFIX string `env:"API_PREFIX,required"`
	APP_ENV    string `env:"APP_ENV" envDefault:"dev"`
}

var AppConfig Config

func LoadAppConfig() {
	_ = godotenv.Load()

	if err := env.Parse(&AppConfig); err != nil {
		logger_configs.Errorw("❌ Failed to load environment variables", "error", err)
		panic("❌ Failed to load environment variables")
	}

	logger_configs.Infof("✅ Environment loaded successfully (%s mode)", AppConfig.APP_ENV)
}
