package main

import (
	"fmt"
	redisDB "url-shortener/internal/adapters/database/redis"
	"url-shortener/internal/adapters/routes"
)

func main() {
	fmt.Println("hello world")

	dbInfra := redisDB.NewInfra()
	router := routes.SetupRouter(dbInfra)
	router.Run(":8080")
}