package ratelimiter

import (
	"fmt"

	"github.com/c-m3-codin/crlim/internal"

	"github.com/c-m3-codin/crlim/config"
)

func NewRateLimiter(cfg *config.Config) *RateLimiter {
	r := &RateLimiter{limits: make(map[string]*TokenBucket)}
	for pattern, limit := range cfg.RateLimits {
		r.limits[pattern] = NewTokenBucket(limit.RequestsPerSecond, limit.BurstSize)
	}
	return r
}

func (r *RateLimiter) Allow(url string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	normalized, err := internal.NormalizeURL(url)
	fmt.Println("notmalised \n ", normalized)
	if err != nil {
		return false
	}
	if pattern, matched := MatchWithWildcard(normalized, r.limits); matched {
		return r.limits[pattern].Allow()
	}
	return true
}
