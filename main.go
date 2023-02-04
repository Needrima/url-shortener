package main

import (
	redisDB "url-shortener/internal/adapters/database/redis"
	"url-shortener/internal/adapters/routes"
	"url-shortener/internal/core/helpers"
)

func main() {
	helpers.InitializeLogger()
	helpers.LogEvent("INFO", "starting server on port 8080")

	dbInfra := redisDB.NewInfra()
	router := routes.SetupRouter(dbInfra)
	router.Run(":8080")
}
