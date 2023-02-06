package ports

import "url-shortener-backend/internal/core/models"

type RedisRepository interface {
	ShortenURL(body models.Request, ip string) map[string]interface{}
	ResolveURL(id, ip string) map[string]interface{}
}
