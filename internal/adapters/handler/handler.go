package handler

import (
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

func (hdl *HTTPHandler) Set(c *gin.Context) {
	hdl.servicePort.Set()
}
func (hdl *HTTPHandler) Get(c *gin.Context) {
	hdl.servicePort.Get()
}