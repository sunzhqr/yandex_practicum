package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func mainPost() {
	postData := `{"login": "some_login@gmail.com", "password": "some_password"}`
	postResponse, err := http.Post("https://stepik.org/users/629159476/profile?auth=login", "application/json", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(postResponse.StatusCode)
	fmt.Println(postResponse.Header.Values("Content-Type"))
	fmt.Println(postResponse.Header.Values("Allow"))
	respBody, err := io.ReadAll(postResponse.Body)
	defer postResponse.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(respBody))
}
