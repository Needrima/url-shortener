package services

import (
	uuid "github.com/satori/go.uuid"
	"time"
	"url-shortener-backend/internal/core/helpers"
	"url-shortener-backend/internal/core/models"
	"url-shortener-backend/internal/ports"
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
			"error": "invalid url",
			"code": 400,
		}
	}
	body.URL = helpers.EnforceHTTP(body.URL)
	body.ExpiriesAt = time.Hour * 24
	if body.CustomID != "" && len(body.CustomID) != 6 { // if user does not specify a custom url id
		return map[string]interface{}{
			"error": "invalid custom id, id must be alphanumeric and six characters long",
			"code": 403,
		}
	}else {
		body.CustomID = uuid.NewV4().String()[:6]
	}

	return s.dbPort.ShortenURL(body, ip)
}
