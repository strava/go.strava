package strava

import (
	"net/http"
	"testing"
	"time"
)

func TestRateLimitUpdating(t *testing.T) {
	var resp http.Response

	resp.StatusCode = 200
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"300,10000"}}
	RateLimiting.updateRateLimits(&resp)

	if !RateLimiting.NextRequestTime.IsZero() {
		t.Errorf("rate limiting didn't set zero time", RateLimiting.NextRequestTime) // non-zero is set only when rate limit is reached
	}

	if RateLimiting.FractionReached() != 0.5 {
		t.Errorf("fraction of rate limit computed incorrectly: %v != %v", RateLimiting.FractionReached(), 0.5)
	}

	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"300,27000"}}
	RateLimiting.updateRateLimits(&resp)

	if !RateLimiting.NextRequestTime.IsZero() {
		t.Errorf("rate limiting didn't set zero time", RateLimiting.NextRequestTime) // non-zero is set only when rate limit is reached
	}

	if RateLimiting.FractionReached() != 0.9 {
		t.Errorf("fraction of rate limit computed incorrectly: %v != %v", RateLimiting.FractionReached(), 0.9)
	}

	// we'll feed it nonsense
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"xxx"}, "X-Ratelimit-Usage": []string{"zzz"}}
	RateLimiting.updateRateLimits(&resp)

	if !RateLimiting.NextRequestTime.IsZero() {
		t.Errorf("nonsense in rate limiting fields should set next reset to zero")
	}

	// rate limit reached - short
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"650,20000"}}
	RateLimiting.updateRateLimits(&resp)

	if RateLimiting.NextRequestTime.IsZero() {
		t.Errorf("reaching rate limit should set non-zero value", RateLimiting.NextRequestTime)
	}

	// rate limit reached - long
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"550,40000"}}
	RateLimiting.updateRateLimits(&resp)

	if RateLimiting.NextRequestTime.IsZero() {
		t.Errorf("reaching rate limit should set non-zero value", RateLimiting.NextRequestTime)
	}
}

func TestNextRateLimit(t *testing.T) {
	// SHORT
	// normal time
	currentTime, err := time.Parse(http.TimeFormat, "Tue, 10 Oct 2013 20:11:05 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	expectedTime, err := time.Parse(http.TimeFormat, "Tue, 10 Oct 2013 20:15:00 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	RateLimiting.NextRequestTime = getNextRateLimitShort(currentTime, currentTime)
	if expectedTime != RateLimiting.NextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, RateLimiting.NextRequestTime)
	}

	// lowest time
	currentTime, err = time.Parse(http.TimeFormat, "Tue, 10 Oct 2013 00:00:01 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	expectedTime, err = time.Parse(http.TimeFormat, "Tue, 10 Oct 2013 00:15:00 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	RateLimiting.NextRequestTime = getNextRateLimitShort(currentTime, currentTime)
	if expectedTime != RateLimiting.NextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, RateLimiting.NextRequestTime)
	}

	// highest time
	currentTime, err = time.Parse(http.TimeFormat, "Tue, 10 Oct 2013 23:59:59 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	expectedTime, err = time.Parse(http.TimeFormat, "Tue, 11 Oct 2013 00:00:00 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	RateLimiting.NextRequestTime = getNextRateLimitShort(currentTime, currentTime)
	if expectedTime != RateLimiting.NextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, RateLimiting.NextRequestTime)
	}

	// LONG
	// normal time
	currentTime, err = time.Parse(http.TimeFormat, "Tue, 10 Oct 2013 20:11:05 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	expectedTime, err = time.Parse(http.TimeFormat, "Tue, 11 Oct 2013 00:00:00 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	RateLimiting.NextRequestTime = getNextRateLimitLong(currentTime, currentTime)
	if expectedTime != RateLimiting.NextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, RateLimiting.NextRequestTime)
	}

	// lowest time
	currentTime, err = time.Parse(http.TimeFormat, "Tue, 10 Oct 2013 00:00:01 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	expectedTime, err = time.Parse(http.TimeFormat, "Tue, 11 Oct 2013 00:00:00 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	RateLimiting.NextRequestTime = getNextRateLimitLong(currentTime, currentTime)
	if expectedTime != RateLimiting.NextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, RateLimiting.NextRequestTime)
	}

	// highest time
	currentTime, err = time.Parse(http.TimeFormat, "Tue, 10 Oct 2013 23:59:59 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	expectedTime, err = time.Parse(http.TimeFormat, "Tue, 11 Oct 2013 00:00:00 GMT")
	if err != nil {
		t.Errorf("error parsing date")
	}
	RateLimiting.NextRequestTime = getNextRateLimitLong(currentTime, currentTime)
	if expectedTime != RateLimiting.NextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, RateLimiting.NextRequestTime)
	}

}

func TestRateLimitChecking(t *testing.T) {
	RateLimiting.NextRequestTime = time.Time{}
	if RateLimiting.Exceeded() == true {
		t.Errorf("rate limiting didn't allow request but should have")
	}

	RateLimiting.NextRequestTime = time.Now().Add(time.Duration(-5 * time.Second))
	if RateLimiting.Exceeded() == true {
		t.Errorf("rate limiting didn't allow request but should have")
	}

	RateLimiting.NextRequestTime = time.Now().Add(time.Duration(5 * time.Second))
	if RateLimiting.Exceeded() == false {
		t.Errorf("rate limiting did allow request but shouldn't have")
	}
}
