package database

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"url-shortener-backend/internal/core/helpers"
	"url-shortener-backend/internal/core/models"

	redis "github.com/go-redis/redis/v8"
)

type RedisInfra struct {
	URLDB    *redis.Client // to store shortened urls based on the id of their short
	IPAddrDB *redis.Client // stores user IP address to check for users usage quota
}

func NewInfra() *RedisInfra {
	URLDB := ConnectToRedis(0)
	IPAddrDB := ConnectToRedis(1)
	return &RedisInfra{
		URLDB:    URLDB,
		IPAddrDB: IPAddrDB,
	}
}

func (r *RedisInfra) ShortenURL(body models.Request, ip string) map[string]interface{} {
	// handle rate limit
	usageTrials, err := r.IPAddrDB.Get(context.TODO(), ip).Result()

	if err != nil {
		switch err {
		case redis.Nil: // user is using the service for the first time or user rate limit has been refreshed
			helpers.LogEvent("INFO", fmt.Sprintf("new user, adding IP address {%v} to database for", ip))

			config := helpers.LoadEnv(".")
			if err := r.IPAddrDB.Set(context.TODO(), ip, config.UsageTrials, time.Minute*30).Err(); err != nil {
				helpers.LogEvent("ERROR", "could not add user's IP address to database:"+err.Error())
				data := map[string]interface{}{
					"error": "something went wrong",
					"code":  500,
				}
				return data
			}

		default:
			helpers.LogEvent("ERROR", "checking database for usage trials err:"+err.Error())
			data := map[string]interface{}{
				"error": "something went wrong",
				"code":  500,
			}
			return data
		}
	} else {
		usageTrialsInt, _ := strconv.Atoi(usageTrials)
		if usageTrialsInt <= 0 { // user exceeded the limit for usage
			helpers.LogEvent("INFO", "user exceeded usage limit")
			limit, _ := r.IPAddrDB.TTL(context.TODO(), ip).Result()
			data := map[string]interface{}{
				"error":             "rate limit exceeded",
				"code":              503,
				"usage_limit_reset": limit / time.Nanosecond / time.Minute, // time to usage limit reset
			}

			return data
		}
	}

	// check if id is already taken
	u, _ := r.URLDB.Get(context.TODO(), body.CustomID).Result()
	if u != "" { // url id is already taken
		helpers.LogEvent("INFO", "url wanted by user is already taken:"+err.Error())
		data := map[string]interface{}{
			"error": "custom id specified for url is already taken",
			"code":  403,
		}
		return data
	}

	// store url in database
	if err := r.URLDB.Set(context.TODO(), body.CustomID, body.URL, body.ExpiriesAt).Err(); err != nil {
		helpers.LogEvent("ERROR", "storing new url in database:"+err.Error())
		data := map[string]interface{}{
			"error": "something went wrong",
			"code":  500,
		}
		return data
	}

	// decrement usage count to track rate limit
	remaining, _ := r.IPAddrDB.Decr(context.TODO(), ip).Result()

	// create response
	config := helpers.LoadEnv(".")
	limit, _ := r.IPAddrDB.TTL(context.TODO(), ip).Result()
	data := map[string]interface{}{
		"url_id": body.CustomID,
		"shortened_url":     fmt.Sprintf("%v/%v", config.Domain, body.CustomID),
		"totalUsageAllowed": 10, // allowed usage per 30 minutes
		"usageRemaining":    remaining,
		"usage_limit_reset": limit / time.Nanosecond / time.Minute, // time to usage limit reset
		"code":              200,
	}

	return data
}

func (r *RedisInfra) ResolveURL(id, ip string) map[string]interface{} {
	fmt.Println("id", id)
	url, err := r.URLDB.Get(context.TODO(), id).Result()
	if err != nil {
		if err == redis.Nil {
			helpers.LogEvent("ERROR", "redis key not found:"+err.Error())
			data := map[string]interface{}{
				"error": "url not found",
				"code":  404,
			}
			return data
		}

		helpers.LogEvent("ERROR", "getting value from db:"+err.Error())
		data := map[string]interface{}{
			"error": "something went wrong",
			"code":  500,
		}
		return data
	}

	data := map[string]interface{}{
		"data": url,
		"code": 200,
	}

	return data
}
