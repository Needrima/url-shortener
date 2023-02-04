package models

import "time"

type Response struct {
	URL            string        `json:"url" binding:"required"`
	CustomShort    string        `json:"custom_short"`
	ExpiriesAt     time.Duration `json:"expires_at"`
	RateLimit      int           `json:"rate_limit"`
	RateLimitReset time.Duration `json:"rate_limit_reset"`
}
