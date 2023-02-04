package ports

type RedisRepository interface {
	Set()
	Get()
}