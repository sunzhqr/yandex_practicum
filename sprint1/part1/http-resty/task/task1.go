package main

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var users []User
	url := "https://jsonplaceholder.typicode.com/users"

	client := resty.New()
	_, err := client.R().
		SetResult(&users).
		Get(url)

	if err != nil {
		panic(err)
	}
	names := make([]string, 0, 10)
	for _, user := range users {
		names = append(names, user.Name)
		fmt.Printf("%v\n", user)
	}
	fmt.Println(strings.Join(names, ", "))
}
