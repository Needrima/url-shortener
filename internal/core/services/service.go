package services

import (
	"url-shortener/internal/core/helpers"
	"url-shortener/internal/core/models"
	"url-shortener/internal/ports"
)

type URLService struct {
	dbPort ports.RedisRepository
}

func NewService(dbPort ports.RedisRepository) *URLService {
	return &URLService{
		dbPort: dbPort,
	}
}

func (s *URLService) ResolveURL(url, ip string) map[string]interface{} {
	return s.dbPort.ResolveURL(url, ip)
}

func (s *URLService) ShortenURL(body models.Request, ip string) map[string]interface{} {
	// url validation
	if err := helpers.ValidateURL(body.URL); err != nil {
		return map[string]interface{}{
			"data":  "invlalid url",
			"error": "",
			"code":  400,
		}
	}
	body.URL = helpers.EnforceHTTP(body.URL)

	return s.dbPort.ShortenURL(body, ip)
}
