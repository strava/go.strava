package strava

import (
	"net/http"
	"testing"
	"time"
)

func TestRateLimitUpdating(t *testing.T) {
	var resp http.Response

	resp.StatusCode = 200
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"50,20000"}}
	updateRateLimits(&resp)

	if !RateLimitingNextRequest.IsZero() {
		t.Errorf("rate limiting didn't set zero time", RateLimitingNextRequest) // non-zero is set only when rate limit is reached
	}

	// we'll feed it nonsense
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"xxx"}, "X-Ratelimit-Usage": []string{"zzz"}}
	updateRateLimits(&resp)

	if !RateLimitingNextRequest.IsZero() {
		t.Errorf("nonsense in rate limiting fields should set next reset to zero")
	}

	// rate limit reached - short
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"650,20000"}}
	updateRateLimits(&resp)

	if RateLimitingNextRequest.IsZero() {
		t.Errorf("reaching rate limit should set non-zero value", RateLimitingNextRequest)
	}

	// rate limit reached - long
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"550,40000"}}
	updateRateLimits(&resp)

	if RateLimitingNextRequest.IsZero() {
		t.Errorf("reaching rate limit should set non-zero value", RateLimitingNextRequest)
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
	nextRequestTime := getNextRateLimitShort(currentTime, currentTime)
	if expectedTime != nextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, nextRequestTime)
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
	nextRequestTime = getNextRateLimitShort(currentTime, currentTime)
	if expectedTime != nextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, nextRequestTime)
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
	nextRequestTime = getNextRateLimitShort(currentTime, currentTime)
	if expectedTime != nextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, nextRequestTime)
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
	nextRequestTime = getNextRateLimitLong(currentTime, currentTime)
	if expectedTime != nextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, nextRequestTime)
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
	nextRequestTime = getNextRateLimitLong(currentTime, currentTime)
	if expectedTime != nextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, nextRequestTime)
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
	nextRequestTime = getNextRateLimitLong(currentTime, currentTime)
	if expectedTime != nextRequestTime {
		t.Errorf("didn't set correct next request time\n%v\n%v", expectedTime, nextRequestTime)
	}

}

func TestRateLimitChecking(t *testing.T) {
	RateLimitingNextRequest = time.Time{}
	if CanDoRequest() == false {
		t.Errorf("rate limiting didn't allow request but should have")
	}

	RateLimitingNextRequest = time.Now().Add(time.Duration(-5 * time.Second))
	if CanDoRequest() == false {
		t.Errorf("rate limiting didn't allow request but should have")
	}

	RateLimitingNextRequest = time.Now().Add(time.Duration(5 * time.Second))
	if CanDoRequest() == true {
		t.Errorf("rate limiting did allow request but shouldn't have")
	}
}
