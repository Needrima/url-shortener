package ports

type URLShortenerService interface {
	Set(data interface{}) (interface{}, error)
	Get()
}
