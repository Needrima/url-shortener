package services

import "url-shortener/internal/ports"

type URLService struct {
	dbPort ports.RedisRepository
}

func NewService(dbPort ports.RedisRepository) *URLService {
	return &URLService{
		dbPort: dbPort,
	}
}

func (s *URLService) Get() {
	s.dbPort.Get()
}

func (s *URLService) Set(data interface{}) (interface{}, error) {
	return s.dbPort.Set(data)
}
