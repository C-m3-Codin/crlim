package crlim

import (
	"fmt"
	"net/http"
)

type RateLimitedClient struct {
	client      *http.Client
	siteLimiter *SiteLimiter
}

func NewRateLimitedClient(policies map[string]RateLimitPolicy) *RateLimitedClient {
	return &RateLimitedClient{
		client:      &http.Client{},
		siteLimiter: NewSiteLimiter(policies),
	}
}

func (c *RateLimitedClient) Do(req *http.Request) (*http.Response, error) {
	if !c.siteLimiter.AllowRequest(req.URL.Host) {
		return nil, fmt.Errorf("rate limit exceeded for %s", req.URL.Host)
	}
	return c.client.Do(req)
}
