package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	if u.FirstName == " " || u.FirstName == "" {
		if u.LastName == " " || u.LastName == "" {
			return ""
		}
		return u.LastName
	}
	if u.LastName == " " || u.LastName == "" {
		return u.FirstName
	}
	return u.FirstName + " " + u.LastName
}

func main() {
	user := User{"", ""}
	fmt.Println(user.FullName(), len(user.FullName()))
}
