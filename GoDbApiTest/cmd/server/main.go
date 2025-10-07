package main

import (
	dbConfig "jackson.com/goApiDb/internal/config"
	albumRouter "jackson.com/goApiDb/internal/router"
)

func main() {
	dbConfig.ConnetDb()
	defer dbConfig.DB.Close()

	router := albumRouter.SetupRouter()
	router.Run("localhost:8080")
}
