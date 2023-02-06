package database

import (
	"url-shortener-backend/internal/core/helpers"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis(dbNo int) *redis.Client {
	config := helpers.LoadEnv(".")

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPass,
		DB:       dbNo,
	})

	return client
}
