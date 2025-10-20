package main

import (
	"jackson.com/libraryapisystem/Configurations/db_configs"
	"jackson.com/libraryapisystem/Configurations/env_configs"
	"jackson.com/libraryapisystem/Configurations/logger_configs"
	"jackson.com/libraryapisystem/Configurations/routes"
)

func main() {

	logger_configs.InitializeLogger(env_configs.AppConfig.APP_ENV)
	defer logger_configs.Sync()

	env_configs.LoadAppConfig()

	logger_configs.Info("🚀 Application started successfully!")

	if err := db_configs.ConnectDatabase(); err != nil {
		logger_configs.Errorf("Database connection failed: %v", err)
		panic(err)
	}

	server := routes.RegisterRoutes()

	logger_configs.Info("🚀 Server running on port %v", env_configs.AppConfig.APP_PORT)
	if err := server.ListenAndServe(); err != nil {
		logger_configs.Errorf("Failed to start server %s", err)
	}

}
