package strava

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// RateLimit is the struct used for the `RateLimiting` global that is
// updated after every request.
type RateLimit struct {
	lock        sync.RWMutex
	RequestTime time.Time
	LimitShort  int
	LimitLong   int
	UsageShort  int
	UsageLong   int
}

// RateLimiting stores rate limit information included in the most recent request.
// Request time will be zero for invalid, or not yet set results.
// Admittedly having a globally updated ratelimit value is a bit clunky. // TODO: fix
var RateLimiting RateLimit

// Exceeded should be called as `strava.RateLimiting.Exceeded() to determine if the most recent
// request exceeded the rate limit
func (rl *RateLimit) Exceeded() bool {
	rl.lock.RLock()
	defer rl.lock.RUnlock()

	if rl.UsageShort >= rl.LimitShort {
		return true
	}

	if rl.UsageLong >= rl.LimitLong {
		return true
	}

	return false
}

// FractionReached returns the current faction of rate used. The greater of the
// short and long term limits. Should be called as `strava.RateLimiting.FractionReached()`
func (rl *RateLimit) FractionReached() float32 {
	rl.lock.RLock()
	defer rl.lock.RUnlock()

	var shortLimitFraction float32 = float32(rl.UsageShort) / float32(rl.LimitShort)
	var longLimitFraction float32 = float32(rl.UsageLong) / float32(rl.LimitLong)

	if shortLimitFraction > longLimitFraction {
		return shortLimitFraction
	} else {
		return longLimitFraction
	}
}

// ignoring error, instead will reset struct to initial values, so rate limiting is ignored
func (rl *RateLimit) updateRateLimits(resp *http.Response) {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	var err error

	if resp.Header.Get("X-Ratelimit-Limit") == "" || resp.Header.Get("X-Ratelimit-Usage") == "" {
		rl.clear()
		return
	}

	s := strings.Split(resp.Header.Get("X-Ratelimit-Limit"), ",")
	if rl.LimitShort, err = strconv.Atoi(s[0]); err != nil {
		rl.clear()
		return
	}
	if rl.LimitLong, err = strconv.Atoi(s[1]); err != nil {
		rl.clear()
		return
	}

	s = strings.Split(resp.Header.Get("X-Ratelimit-Usage"), ",")
	if rl.UsageShort, err = strconv.Atoi(s[0]); err != nil {
		rl.clear()
		return
	}

	if rl.UsageLong, err = strconv.Atoi(s[1]); err != nil {
		rl.clear()
		return
	}

	rl.RequestTime = time.Now()
	return
}

func (rl *RateLimit) clear() {
	rl.RequestTime = time.Time{}
	rl.LimitShort = 0
	rl.LimitLong = 0
	rl.UsageShort = 0
	rl.UsageLong = 0
}
