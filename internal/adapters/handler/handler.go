package handler

import (
	"log"
	"url-shortener/internal/core/models"
	"url-shortener/internal/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	servicePort ports.URLShortenerService
}

func NewHandler(servicePort ports.URLShortenerService) *HTTPHandler {
	return &HTTPHandler{
		servicePort: servicePort,
	}
}

func (hdl *HTTPHandler) ShortenURL(c *gin.Context) {
	body := models.Request{}

	if err := c.BindJSON(&body); err != nil {
		log.Println("error reading request body:", err.Error())
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}

	res, err := hdl.servicePort.ShortenURL(body)
	if err != nil {
		log.Println("error shortening URL:", err.Error())
		c.JSON(500, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(200, res)
}
func (hdl *HTTPHandler) ResolveURL(c *gin.Context) {
	hdl.servicePort.ResolveURL()
}
