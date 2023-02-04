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

func (s *URLService) ResolveURL() {
	s.dbPort.ResolveURL()
}

func (s *URLService) ShortenURL(data models.Request) (interface{}, error) {
	if err := helpers.ValidateURL(data.URL); err != nil {
		return nil, err
	}
	data.URL = helpers.EnforceHTTP(data.URL)

	return s.dbPort.ShortenURL(data)
}
