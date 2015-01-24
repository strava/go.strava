package strava

var RateLimitLast RateLimit

type RateLimit struct {
	LimitShort int
	LimitLong  int
	UsageShort int
	UsageLong  int
}

func RateLimitReached() bool {
	if RateLimitLast.UsageShort == 0 { // no need to check both values
		return false
	} else if RateLimitLast.UsageShort >= RateLimitLast.LimitShort || RateLimitLast.UsageLong >= RateLimitLast.LimitLong {
		return true
	} else {
		return false
	}
}
