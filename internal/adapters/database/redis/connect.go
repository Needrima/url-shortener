package database

import (
	"url-shortener/internal/core/helpers"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis(dbNo int) *redis.Client {
	config, err := helpers.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPass,
		DB:       dbNo,
	})

	return client
}
