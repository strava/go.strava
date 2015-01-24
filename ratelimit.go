package strava

import (
	"errors"
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

func updateRateLimits(resp *http.Response) error {
	var err error

	if resp.Header.Get("X-Ratelimit-Limit") == "" || resp.Header.Get("X-Ratelimit-Usage") == "" {
		return errors.New("ratelimit headers not found")
	}

	s := strings.Split(resp.Header.Get("X-Ratelimit-Limit"), ",")
	if RateLimitLast.LimitShort, err = strconv.Atoi(s[0]); err != nil {
		return err
	}
	if RateLimitLast.LimitLong, err = strconv.Atoi(s[1]); err != nil {
		return err
	}

	s = strings.Split(resp.Header.Get("X-Ratelimit-Usage"), ",")
	if RateLimitLast.UsageShort, err = strconv.Atoi(s[0]); err != nil {
		return err
	}
	if RateLimitLast.UsageLong, err = strconv.Atoi(s[1]); err != nil {
		return err
	}

	RateLimitLast.LastRequestTime = time.Now()

	return nil
}
