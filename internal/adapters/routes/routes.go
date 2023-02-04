package routes

import (
	"url-shortener/internal/adapters/handler"
	"url-shortener/internal/core/middleware"
	"url-shortener/internal/core/services"
	"url-shortener/internal/ports"

	"github.com/gin-gonic/gin"
)

func SetupRouter(dbRepository ports.RedisRepository) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware)

	service := services.NewService(dbRepository)

	handler := handler.NewHandler(service)

	routerGroup := router.Group("/api")
	{
		routerGroup.POST("/shorten_url", handler.ShortenURL)
		routerGroup.GET("/:url", handler.ResolveURL)
	}

	return router
}
