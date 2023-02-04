package ports

type RedisRepository interface {
	Set(data interface{}) (interface{}, error)
	Get()
}
