package ratelimiter

import (
	"net/http"
)

type RateLimitedClient struct {
	client      *http.Client
	rateLimiter *RateLimiter
}

func NewRateLimitedClient(rateLimiter *RateLimiter) *RateLimitedClient {
	return &RateLimitedClient{
		client:      &http.Client{},
		rateLimiter: rateLimiter,
	}
}

func (c *RateLimitedClient) Do(req *http.Request) (*http.Response, error) {
	if !c.rateLimiter.Allow(req.URL.String()) {
		return nil, http.ErrHandlerTimeout
	}
	return c.client.Do(req)
}
