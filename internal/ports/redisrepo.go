package ports

import "url-shortener/internal/core/models"

type RedisRepository interface {
	ShortenURL(data models.Request) (interface{}, error)
	ResolveURL()
}
