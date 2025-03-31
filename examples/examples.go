package main

import (
	"fmt"
	"net/http"

	"github.com/c-m3-codin/crlim"
)

func main() {
	policies := map[string]crlim.RateLimitPolicy{
		"api.example.com": {RequestsPerSecond: 10, BurstSize: 5},
		"other.com":       {RequestsPerSecond: 5, BurstSize: 2},
	}

	client := crlim.NewRateLimitedClient(policies)

	req, _ := http.NewRequest("GET", "https://api.example.com/data", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", resp.Status)
	}
}
