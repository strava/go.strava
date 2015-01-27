package strava

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

var RateLimitingNextRequest time.Time

func CanDoRequest() bool {
	if time.Now().After(RateLimitingNextRequest) {
		return true
	} else {
		return false
	}
}

// ignoring error, instead will reset struct to initial values, so rate limiting is ignored
func updateRateLimits(resp *http.Response) {
	if resp.Header.Get("X-Ratelimit-Limit") == "" || resp.Header.Get("X-Ratelimit-Usage") == "" {
		RateLimitingNextRequest = time.Time{}
		return
	}

	s := strings.Split(resp.Header.Get("X-Ratelimit-Limit"), ",")
	limitShort, err := strconv.Atoi(s[0])
	if err != nil {
		RateLimitingNextRequest = time.Time{}
		return
	}
	limitLong, err := strconv.Atoi(s[1])
	if err != nil {
		RateLimitingNextRequest = time.Time{}
		return
	}

	s = strings.Split(resp.Header.Get("X-Ratelimit-Usage"), ",")
	usageShort, err := strconv.Atoi(s[0])
	if err != nil {
		RateLimitingNextRequest = time.Time{}
		return
	}
	usageLong, err := strconv.Atoi(s[1])
	if err != nil {
		RateLimitingNextRequest = time.Time{}
		return
	}

	currentServerTime, err := time.Parse(http.TimeFormat, resp.Header.Get("Date"))
	if err != nil {
		RateLimitingNextRequest = time.Time{}
	} else if usageShort >= limitShort {
		RateLimitingNextRequest = getNextRateLimitShort(time.Now(), currentServerTime)
	} else if usageLong >= limitLong {
		RateLimitingNextRequest = getNextRateLimitLong(time.Now(), currentServerTime)
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
