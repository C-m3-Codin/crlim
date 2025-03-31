package main

import (
	"log"
	"net/http"
	"time"

	"github.com/c-m3-codin/crlim/config"
	"github.com/c-m3-codin/crlim/ratelimiter"
)

func main() {
	// Load rate limit config
	cfg, err := config.LoadConfig("config.json") // or "config.yaml"
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Initialize rate limiter
	limiter := ratelimiter.NewRateLimiter(cfg)

	// Create a rate-limited HTTP client
	client := ratelimiter.NewRateLimitedClient(limiter)

	// Example requests
	urls := []string{
		"https://api.example.com/v1/users",
		"https://api.example.com/v1/orders",
		"https://sub.example.com/data",
	}

	for _, url := range urls {
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)

		if err != nil {
			log.Printf("Request blocked by rate limiter for %s: %v\n", url, err)
		} else {
			log.Printf("Request successful for %s: %s\n", url, resp.Status)
		}

		time.Sleep(500 * time.Millisecond) // Simulate request intervals
	}
}
