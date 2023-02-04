package routes

import (
	"url-shortener/internal/ports"
	"url-shortener/internal/core/services"
	"url-shortener/internal/adapters/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(dbRepository ports.RedisRepository) *gin.Engine {
	router := gin.Default()
	service := services.NewService(dbRepository)
	handler := handler.NewHandler(service)

	routerGroup := router.Group("/api")
	{
		routerGroup.POST("/set", handler.Set)
		routerGroup.GET("/get/:key", handler.Get)
	}

	return router
}