package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func mainAuth() {
	// create new client
	client := resty.New()
	// client.R() creates new Request(and returns it like *Request)
	// SetAuthToken() sets token
	// Get() is standard http.Get()
	resp, err := client.R().Get("https://github.com/go-resty/resty")
	if err != nil {
		panic(err)
	}
	fmt.Println("Explore Response object:")
	fmt.Println("Error:", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status:", resp.Status())
	fmt.Println("Time:", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	// fmt.Println("Body       :\n", resp)
	fmt.Println("----")
}
