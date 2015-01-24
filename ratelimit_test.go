package strava

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRateLimitUpdating(t *testing.T) {
	var resp http.Response

	resp.StatusCode = 200
	resp.Header = http.Header{"X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"50,20000"}}
	updateRateLimits(&resp)

	// not comparing time
	expected := RateLimit{LimitShort: 600, LimitLong: 30000, UsageShort: 50, UsageLong: 20000, LastRequestTime: RateLimitLast.LastRequestTime}

	if !reflect.DeepEqual(RateLimitLast, expected) {
		t.Errorf("should match\n%v\n%v", RateLimitLast, expected)
	}

	if RateLimitLast.LastRequestTime.IsZero() {
		t.Errorf("rate limiting didn't set non-zero time")
	}

	// we'll feed it nonsense
	resp.Header = http.Header{"X-Ratelimit-Limit": []string{"xxx"}, "X-Ratelimit-Usage": []string{"zzz"}}
	updateRateLimits(&resp)

	expected = RateLimit{}

	if !reflect.DeepEqual(RateLimitLast, expected) {
		t.Errorf("should match\n%v\n%v", RateLimitLast, expected)
	}

	// fill it with real values and then test missing headers
	resp.Header = http.Header{"X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"50,20000"}}
	updateRateLimits(&resp)
	resp.Header = http.Header{}
	updateRateLimits(&resp)

	if !reflect.DeepEqual(RateLimitLast, expected) {
		t.Errorf("should match\n%v\n%v", RateLimitLast, expected)
	}
}

func TestRateLimitChecking(t *testing.T) {
	var resp http.Response

	resp.StatusCode = 200
	resp.Header = http.Header{"X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"50,20000"}}
	updateRateLimits(&resp)

	if RateLimitReachedDuringLast(60) {
		t.Errorf("rate limit reached when it shouldn't have been")
	}

	resp.Header = http.Header{"X-Ratelimit-Limit": []string{"600,30000"}, "X-Ratelimit-Usage": []string{"650,20000"}}
	updateRateLimits(&resp)

	if !RateLimitReachedDuringLast(60) {
		t.Errorf("rate limit not reached when it should have been")
	}

	// if rate limit was triggered before "future", it must be expired by now
	if RateLimitReachedDuringLast(-5) {
		t.Errorf("rate limit not reached when it should have been")
	}

	resp.Header = http.Header{"X-Ratelimit-Limit": []string{"nonsense"}, "X-Ratelimit-Usage": []string{"EPO"}}
	updateRateLimits(&resp)

	if RateLimitReachedDuringLast(60) {
		t.Errorf("rate limit not reached when it should have been")
	}
}
