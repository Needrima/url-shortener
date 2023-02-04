package database

import "github.com/go-redis/redis/v8"

func ConnectToRedis(dbNo int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "db:6379",
		Password: "",
		DB: dbNo,
	})

	return client
}