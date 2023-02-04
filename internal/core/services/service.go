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

func (s *URLService)Get() {
	s.dbPort.Get()
}

func (s *URLService)Set() {
	s.dbPort.Set()
}