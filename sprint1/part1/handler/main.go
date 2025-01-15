package main

import (
	"log"
	"net/http"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
}

func main() {
	users := make(map[string]User)
	user1 := User{
		"u1", "Sanzhar", "Myrzash",
	}
	user2 := User{
		"u2", "Bauyrzhan", "Altai",
	}
	users["u1"] = user1
	users["u2"] = user2

	http.Handle("/users", UserViewHandler(users))
	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/url", URLHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
