package ratelimiter

import (
	"time"
)

func NewTokenBucket(ratePerSecond float64, capacity int) *TokenBucket {
	return &TokenBucket{
		tokens:         float64(capacity),
		capacity:       float64(capacity),
		ratePerSecond:  ratePerSecond,
		lastRefillTime: time.Now(),
	}
}

func (t *TokenBucket) Allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	now := time.Now()
	duration := now.Sub(t.lastRefillTime).Seconds()
	t.tokens += duration * t.ratePerSecond
	if t.tokens > t.capacity {
		t.tokens = t.capacity
	}
	t.lastRefillTime = now

	if t.tokens >= 1 {
		t.tokens--
		return true
	}
	return false
}
