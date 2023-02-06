package main

import (
	redisDB "url-shortener-backend/internal/adapters/database/redis"
	"url-shortener-backend/internal/adapters/routes"
	"url-shortener-backend/internal/core/helpers"
)

func main() {
	helpers.InitializeLogger()
	helpers.LogEvent("INFO", "starting server on port 8080")

	dbInfra := redisDB.NewInfra()
	router := routes.SetupRouter(dbInfra)
	config := helpers.LoadEnv(".")
	router.Run(":" + config.Port)
}
