package ports

import "url-shortener/internal/core/models"

type URLShortenerService interface {
	ShortenURL(data models.Request) (interface{}, error)
	ResolveURL()
}
