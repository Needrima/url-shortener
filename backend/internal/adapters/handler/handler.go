package handler

import (
	"log"
	"url-shortener-backend/internal/core/models"
	"url-shortener-backend/internal/ports"

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
			"error": "bad request, check that field are correct",
			"reasons": []string{
				"user specified a custom id with non alphanumeric characters",
				"user specified id not up to six characters",
			},
		})
		return
	}
	ip := c.RemoteIP()

	res := hdl.servicePort.ShortenURL(body, ip)

	c.JSON(res["code"].(int), res)
}
func (hdl *HTTPHandler) ResolveURL(c *gin.Context) {
	id, ip := c.Param("id"), c.RemoteIP()
	res := hdl.servicePort.ResolveURL(id, ip)

	c.JSON(res["code"].(int), res)
}
