package ports

import "url-shortener/internal/core/models"

type RedisRepository interface {
	ShortenURL(body models.Request, ip string) map[string]interface{}
	ResolveURL(url, ip string) map[string]interface{}
}
