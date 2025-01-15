package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	client.
		SetRetryCount(3).
		SetRetryWaitTime(30 * time.Second).
		SetRetryMaxWaitTime(90 * time.Second)

	req := client.R()
	//	Way 1

	/*req.SetHeader("Content-Type", "application/json")
	req.SetBody(`{"title": "foo", "body": "bar", "userId": "7"}`)*/

	// Way 2 ---With using map in SetBody we can omit SetHeader---
	// because it's already created Content-Type header if body provides with map

	req.SetBody(map[string]interface{}{
		"title":  "Jah Khalib",
		"body":   "Your sleepy eyes",
		"userId": "10",
	})

	resp, err := req.
		Post("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Header().Get("Content-Type"))
}
