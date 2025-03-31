package crlim

import (
	"sync"

	"golang.org/x/time/rate"
)

type RateLimitPolicy struct {
	RequestsPerSecond int
	BurstSize         int
}

type SiteLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
}

func NewSiteLimiter(policies map[string]RateLimitPolicy) *SiteLimiter {
	limiters := make(map[string]*rate.Limiter)
	for site, policy := range policies {
		limiters[site] = rate.NewLimiter(rate.Limit(policy.RequestsPerSecond), policy.BurstSize)
	}
	return &SiteLimiter{limiters: limiters}
}

func (s *SiteLimiter) AllowRequest(host string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if limiter, exists := s.limiters[host]; exists {
		return limiter.Allow()
	}
	return true // If no policy exists, allow by default
}
