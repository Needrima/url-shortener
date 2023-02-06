package models

import (
	"time"
)

type Request struct {
	URL        string        `json:"url" binding:"required"` // url to be shortened
	CustomID   string        `json:"custom_id"`              // user defined custom url id
	ExpiriesAt time.Duration `json:"expires_at"`
}
