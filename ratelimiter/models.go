package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	limits map[string]*TokenBucket
	mu     sync.Mutex
}

type TokenBucket struct {
	tokens         float64
	capacity       float64
	ratePerSecond  float64
	lastRefillTime time.Time
	mu             sync.Mutex
}
