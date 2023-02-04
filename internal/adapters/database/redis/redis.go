package database

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisInfra struct{
	client *redis.Client
}

func NewInfra(client *redis.Client) *RedisInfra {
	return &RedisInfra{
		client: client,
	}
}

func (r *RedisInfra) Set() {
	r.client.Set(context.TODO(), "key", map[string]interface{}{}, time.Minute)
}

func (r *RedisInfra) Get() {
	r.client.Get(context.TODO(), "key")
}