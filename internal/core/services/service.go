package services

import (
	uuid "github.com/satori/go.uuid"
	"time"
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

func (s *URLService) ResolveURL(id, ip string) map[string]interface{} {
	return s.dbPort.ResolveURL(id, ip)
}

func (s *URLService) ShortenURL(body models.Request, ip string) map[string]interface{} {
	// url validation
	if err := helpers.ValidateURL(body.URL); err != nil {
		return map[string]interface{}{
			"data": "invlalid url",
			"code": 400,
		}
	}
	body.URL = helpers.EnforceHTTP(body.URL)
	body.ExpiriesAt = time.Hour * 24
	if body.CustomID == "" { // if user does not specify a custom url id
		body.CustomID = uuid.NewV4().String()[:6]
	}

	return s.dbPort.ShortenURL(body, ip)
}
