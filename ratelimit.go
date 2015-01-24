package strava

import (
	"time"
)

var RateLimitLast RateLimit

type RateLimit struct {
	LastRequestTime time.Time
	LimitShort      int
	LimitLong       int
	UsageShort      int
	UsageLong       int
}

// returns true if rate limit was reached during last X seconds
func RateLimitReachedDuringLast(seconds int64) bool {
	if RateLimitLast.LastRequestTime.IsZero() {
		// no idea, so we should try
		return false
	} else if RateLimitLast.LastRequestTime.Unix() < (time.Now().Unix() - seconds) {
		// last request was some time ago, so we should try
		return false
	} else if RateLimitLast.UsageShort < RateLimitLast.LimitShort && RateLimitLast.UsageLong < RateLimitLast.LimitLong {
		// limit not reached
		return false
	} else {
		return true
	}
}
