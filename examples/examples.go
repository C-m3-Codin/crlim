package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/c-m3-codin/crlim"
)

func main() {
	config, err := crlim.LoadConfig("./config_examples/config.json")

	client := crlim.NewRateLimitedClient(config.RateLimits)

	req, _ := http.NewRequest("GET", "https://www.google.com", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", resp.Status)
	}
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func() {
			_, err := client.Do(req)
			if err != nil {
				fmt.Printf("Request no:%d Error:%s \n", i, err)
			} else {
				fmt.Printf("Response no:%d Response status:%s \n", i, resp.Status)
			}
			wg.Done()

		}()
	}
	wg.Wait()
}
