package database

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"url-shortener/internal/core/helpers"
	"url-shortener/internal/core/models"

	redis "github.com/go-redis/redis/v8"
)

type RedisInfra struct {
	URLclient    *redis.Client // to store shortened URLs
	IPAddrClient *redis.Client // stores user IP address to check for users usage quota
}

func NewInfra() *RedisInfra {
	URLclient := ConnectToRedis(0)
	IPAddrClient := ConnectToRedis(1)
	return &RedisInfra{
		URLclient:    URLclient,
		IPAddrClient: IPAddrClient,
	}
}

func (r *RedisInfra) ShortenURL(body models.Request, ip string) map[string]interface{} {
	usageTrials, err := r.IPAddrClient.Get(context.TODO(), ip).Result()

	if err != nil {
		switch err {
		case redis.Nil: // no records found for user IP
			helpers.LogEvent("INFO", fmt.Sprintf("new user, adding IP address {%v} to database for", ip))

			config := helpers.LoadEnv(".")
			if err := r.IPAddrClient.Set(context.TODO(), ip, config.UsageTrials, time.Minute*30).Err(); err != nil {
				helpers.LogEvent("ERROR", "could not add user's IP address to database:"+err.Error())
				data := map[string]interface{}{
					"data":  "",
					"error": "something went wrong",
					"code":  500,
				}
				return data
			}

		default:
			helpers.LogEvent("ERROR", "checking database for usage trials err:"+err.Error())
			data := map[string]interface{}{
				"data":  "",
				"error": "something went wrong",
				"code":  500,
			}
			return data
		}
	} else {
		usageTrialsInt, _ := strconv.Atoi(usageTrials)
		if usageTrialsInt <= 0 { // user exceeded the limit for usage
			helpers.LogEvent("INFO", "user exceede rate limit for")
			limit, _ := r.IPAddrClient.TTL(context.TODO(), ip).Result()
			data := map[string]interface{}{
				"data":             "",
				"error":            "rate limit exceeded",
				"code":             503,
				"rate_limit_reset": limit / time.Nanosecond / time.Minute, // time to usage limit reset
			}

			return data
		}
	}

	r.IPAddrClient.Decr(context.TODO(), ip) // decrement usage count to track rate limit
	
	return map[string]interface{}{}
}

func (r *RedisInfra) ResolveURL(url, ip string) map[string]interface{} {
	short, err := r.URLclient.Get(context.TODO(), url).Result()
	if err != nil {
		if err == redis.Nil {
			helpers.LogEvent("ERROR", "redis key not found:"+err.Error())
			data := map[string]interface{}{
				"data":  "",
				"error": "url not found",
				"code":  404,
			}
			return data
		}

		helpers.LogEvent("ERROR", "getting value from db:"+err.Error())
		data := map[string]interface{}{
			"data":  "",
			"error": "something went wrong",
			"code":  500,
		}
		return data
	}

	r.IPAddrClient.Decr(context.TODO(), ip)

	data := map[string]interface{}{
		"data":  short,
		"error": "",
		"code":  200,
	}

	return data
}
