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
	ip := c.RemoteIP()

	res := hdl.servicePort.ShortenURL(body, ip)

	c.JSON(res["code"].(int), res)
}
func (hdl *HTTPHandler) ResolveURL(c *gin.Context) {
	url, ip := c.Param("url"), c.RemoteIP()
	res := hdl.servicePort.ResolveURL(url, ip)

	c.JSON(res["code"].(int), res)
}
