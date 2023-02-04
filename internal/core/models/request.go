package models

import (
	"time"
)

type Request struct {
	URL         string        `json:"url" binding:"required"`
	CustomShort string        `json:"custom_short"`
	ExpiriesAt  time.Duration `json:"expires_at"`
}
