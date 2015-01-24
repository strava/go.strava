package strava

import (
	"net/http"
	"strconv"
	"strings"
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

// ignoring error, inster will reset struct to initial values, so rate limiting is ignored
func updateRateLimits(resp *http.Response) {
	var err error

	if resp.Header.Get("X-Ratelimit-Limit") == "" || resp.Header.Get("X-Ratelimit-Usage") == "" {
		RateLimitLast = RateLimit{}
		return
	}

	s := strings.Split(resp.Header.Get("X-Ratelimit-Limit"), ",")
	if RateLimitLast.LimitShort, err = strconv.Atoi(s[0]); err != nil {
		RateLimitLast = RateLimit{}
		return
	}
	if RateLimitLast.LimitLong, err = strconv.Atoi(s[1]); err != nil {
		RateLimitLast = RateLimit{}
		return
	}

	s = strings.Split(resp.Header.Get("X-Ratelimit-Usage"), ",")
	if RateLimitLast.UsageShort, err = strconv.Atoi(s[0]); err != nil {
		RateLimitLast = RateLimit{}
		return
	}
	if RateLimitLast.UsageLong, err = strconv.Atoi(s[1]); err != nil {
		RateLimitLast = RateLimit{}
		return
	}

	RateLimitLast.LastRequestTime = time.Now()
}
