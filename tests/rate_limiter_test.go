package crlim

import (
	"testing"
	"time"

	"github.com/c-m3-codin/crlim"
)

func TestRateLimiting(t *testing.T) {
	policies := map[string]crlim.RateLimitPolicy{
		"example.com": {RequestsPerSecond: 1, BurstSize: 1},
	}

	limiter := crlim.NewSiteLimiter(policies)
	if !limiter.AllowRequest("example.com") {
		t.Error("Expected request to be allowed initially")
	}

	if limiter.AllowRequest("example.com") {
		t.Error("Expected request to be blocked due to rate limit")
	}

	time.Sleep(time.Second) // Allow refill
	if !limiter.AllowRequest("example.com") {
		t.Error("Expected request to be allowed after refill")
	}
}
