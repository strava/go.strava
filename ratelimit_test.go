package strava

import (
	"net/http"
	"testing"
)

func TestRateLimitUpdating(t *testing.T) {
	var resp http.Response

	resp.StatusCode = 200
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"300,10000"}}
	RateLimiting.updateRateLimits(&resp)

	if RateLimiting.RequestTime.IsZero() {
		t.Errorf("rate limiting should set request time")
	}

	if v := RateLimiting.FractionReached(); v != 0.5 {
		t.Errorf("fraction of rate limit computed incorrectly, got %v", v)
	}

	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"300,27000"}}
	RateLimiting.updateRateLimits(&resp)

	if RateLimiting.RequestTime.IsZero() {
		t.Errorf("rate limiting should set request time")
	}

	if v := RateLimiting.FractionReached(); v != 0.9 {
		t.Errorf("fraction of rate limit computed incorrectly, got %v", v)
	}

	// we'll feed it nonsense
	resp.Header = http.Header{"Date": []string{"Tue, 10 Oct 2013 20:11:05 GMT"}, "X-Ratelimit-Limit": []string{"xxx"}, "X-Ratelimit-Usage": []string{"zzz"}}
	RateLimiting.updateRateLimits(&resp)

	if !RateLimiting.RequestTime.IsZero() {
		t.Errorf("nonsense in rate limiting fields should set next reset to zero")
	}
}

func TestRateLimitExceeded(t *testing.T) {
	RateLimiting.LimitLong = 1
	RateLimiting.UsageLong = 0

	RateLimiting.LimitShort = 100
	RateLimiting.UsageShort = 200

	if RateLimiting.Exceeded() != true {
		t.Errorf("should have exceeded rate limit")
	}

	RateLimiting.LimitShort = 200
	RateLimiting.UsageShort = 100

	if RateLimiting.Exceeded() == true {
		t.Errorf("should not have exceeded rate limit")
	}

	RateLimiting.LimitShort = 1
	RateLimiting.UsageShort = 0
	RateLimiting.LimitLong = 100
	RateLimiting.UsageLong = 200

	if RateLimiting.Exceeded() != true {
		t.Errorf("should have exceeded rate limit")
	}

	RateLimiting.LimitLong = 200
	RateLimiting.UsageLong = 100

	if RateLimiting.Exceeded() == true {
		t.Errorf("should not have exceeded rate limit")
	}
}
