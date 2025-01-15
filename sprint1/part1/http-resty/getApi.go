package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type MyApiError struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"body"`
}

func mainGet() {
	client := resty.New()

	var responseError MyApiError
	var post Post

	_, err := client.R().
		SetError(&responseError).
		SetResult(&post).
		SetPathParams(map[string]string{
			"postID": "2",
		}).
		Get("https://jsonplaceholder.typicode.com/posts/{postID}")

	if err != nil {
		fmt.Println(responseError)
		panic(err)
		return
	}

	fmt.Println(post)
}
