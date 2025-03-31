package ratelimiter

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/c-m3-codin/crlim/config"
)

// TestRateLimiting ensures rate limits work correctly
func TestRateLimiting(t *testing.T) {
	// Define rate limit config
	cfg := &config.Config{
		RateLimits: map[string]config.RateLimitConfig{
			"example.com/api/*":  {RequestsPerSecond: 2, BurstSize: 1},
			"example.com/data/*": {RequestsPerSecond: 1, BurstSize: 1},
		},
	}

	// Initialize rate limiter
	limiter := NewRateLimiter(cfg)

	// First request (should be allowed)
	if !limiter.Allow("https://example.com/api/users") {
		t.Error("Expected request to be allowed initially")
	}
	time.Sleep(time.Second)

	// Second request (should be allowed, burst of 1)
	if !limiter.Allow("https://example.com/api/orders") {
		t.Error("Expected second request within limit")
	}

	// Third request (should be blocked due to rate limit)

	blocked_list := make([]bool, 5)
	var wg sync.WaitGroup
	for i := range 5 {
		go func(i int) {
			blocked_list[i] = limiter.Allow("https://example.com/api/products")
			wg.Done()
		}(i)
		wg.Add(1)

	}
	wg.Wait()

	unblocked := true
	for i := range 5 {
		fmt.Println(blocked_list[i])
		unblocked = unblocked && blocked_list[i]
	}

	if unblocked {
		t.Error("Expected request to be blocked due to rate limit")
	}

	// Wait for token refill
	time.Sleep(time.Second)

	// Request after refill (should be allowed again)
	if !limiter.Allow("https://example.com/api/users") {
		t.Error("Expected request to be allowed after refill")
	}
}

// TestWildcardMatching ensures wildcard-based limits work correctly
func TestWildcardMatching(t *testing.T) {
	cfg := &config.Config{
		RateLimits: map[string]config.RateLimitConfig{
			"example.com/api/*":  {RequestsPerSecond: 3, BurstSize: 1},
			"api.example.com/*":  {RequestsPerSecond: 1, BurstSize: 1},
			"example.com/static": {RequestsPerSecond: 1, BurstSize: 1},
		},
	}

	limiter := NewRateLimiter(cfg)

	// Wildcard should match
	if !limiter.Allow("https://example.com/api/users") {
		t.Error("Expected wildcard match for example.com/api/users")
	}
	time.Sleep(time.Second)
	if !limiter.Allow("https://example.com/api/orders") {
		t.Error("Expected wildcard match for example.com/api/orders")
	}
	time.Sleep(time.Second)
	if !limiter.Allow("https://api.example.com/dashboard") {
		t.Error("Expected wildcard match for api.example.com/dashboard")
	}
	time.Sleep(time.Second)

	// Exact match should also work
	if !limiter.Allow("https://example.com/static") {
		t.Error("Expected exact match for example.com/static")
	}
}
