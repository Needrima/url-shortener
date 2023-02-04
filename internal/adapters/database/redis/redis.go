package database

import (
	"context"
	"errors"
	"log"
	"url-shortener/internal/core/models"

	redis "github.com/go-redis/redis/v8"
)

type RedisInfra struct {
	client *redis.Client
}

func NewInfra() *RedisInfra {
	client := ConnectToRedis(0)
	return &RedisInfra{
		client: client,
	}
}

func (r *RedisInfra) Set(data interface{}) (interface{}, error) {
	if err := r.client.Set(context.TODO(), "key", data, 0).Err(); err != nil {
		log.Println("error storing URL in redis cache:", err.Error())
		return nil, errors.New("something went wrong")
	}

	return models.Response{}, nil
}

func (r *RedisInfra) Get() {
	r.client.Get(context.TODO(), "key").Result()
}
