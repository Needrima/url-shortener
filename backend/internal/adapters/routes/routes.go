package routes

import (
	"url-shortener-backend/internal/adapters/handler"
	"url-shortener-backend/internal/core/middleware"
	"url-shortener-backend/internal/core/services"
	"url-shortener-backend/internal/ports"

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
		routerGroup.GET("/:id", handler.ResolveURL)
	}

	return router
}
