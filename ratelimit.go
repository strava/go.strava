package strava

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RateLimit struct {
	NextRequestTime time.Time
	LimitShort      int
	LimitLong       int
	UsageShort      int
	UsageLong       int
}

var RateLimiting RateLimit

func (ratelimit *RateLimit) Exceeded() bool {
	if time.Now().After(ratelimit.NextRequestTime) {
		return false
	} else {
		return true
	}
}

func (ratelimit *RateLimit) FractionReached() float32 {
	var shortLimitReached float32 = float32(ratelimit.UsageShort) / float32(ratelimit.LimitShort)
	var longLimitReached float32 = float32(ratelimit.UsageLong) / float32(ratelimit.LimitLong)

	if shortLimitReached > longLimitReached {
		return shortLimitReached
	} else {
		return longLimitReached
	}
}

// ignoring error, instead will reset struct to initial values, so rate limiting is ignored
func (ratelimit *RateLimit) updateRateLimits(resp *http.Response) {
	var err error

	if resp.Header.Get("X-Ratelimit-Limit") == "" || resp.Header.Get("X-Ratelimit-Usage") == "" {
		ratelimit = &RateLimit{}
		return
	}

	s := strings.Split(resp.Header.Get("X-Ratelimit-Limit"), ",")
	if ratelimit.LimitShort, err = strconv.Atoi(s[0]); err != nil {
		ratelimit = &RateLimit{}
		return
	}
	if ratelimit.LimitLong, err = strconv.Atoi(s[1]); err != nil {
		ratelimit = &RateLimit{}
		return
	}

	s = strings.Split(resp.Header.Get("X-Ratelimit-Usage"), ",")
	if ratelimit.UsageShort, err = strconv.Atoi(s[0]); err != nil {
		ratelimit = &RateLimit{}
		return
	}

	if ratelimit.UsageLong, err = strconv.Atoi(s[1]); err != nil {
		ratelimit = &RateLimit{}
		return
	}

	currentServerTime, err := time.Parse(http.TimeFormat, resp.Header.Get("Date"))
	if err != nil {
		ratelimit.NextRequestTime = time.Time{}
	} else if ratelimit.UsageShort >= ratelimit.LimitShort {
		ratelimit.NextRequestTime = getNextRateLimitShort(time.Now(), currentServerTime)
	} else if ratelimit.UsageLong >= ratelimit.LimitLong {
		ratelimit.NextRequestTime = getNextRateLimitLong(time.Now(), currentServerTime)
	}
}

func getNextRateLimitShort(localTime time.Time, serverTime time.Time) time.Time {
	var timeRemaining time.Duration = time.Second * time.Duration(900-(serverTime.Minute()*60+serverTime.Second())%900)
	return localTime.Add(timeRemaining)
}

func getNextRateLimitLong(localTime time.Time, serverTime time.Time) time.Time {
	var timeRemaining time.Duration = time.Second * time.Duration(86400-(serverTime.Hour()*3600+serverTime.Minute()*60+serverTime.Second())%86400)
	return localTime.Add(timeRemaining)
}
