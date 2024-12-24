package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)

var pl = fmt.Println

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			pl(req)
			return nil
		},
		Jar: jar,
	}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/status", nil)
	if err != nil {
		pl(err)
		return
	}
	request.AddCookie(&http.Cookie{
		Name:   "access_token",
		Value:  "050422",
		MaxAge: 60,
	})

	request.Header.Set("MyHeader", "Hello")
	request.Header.Add("MyHeader", "World")

	response, err := client.Do(request)
	if err != nil {
		pl(err)
		return
	}
	fmt.Println("---Response Headers---")
	for headerName, headerValues := range response.Header {
		fmt.Printf("%s: %v\r\n", headerName, headerValues)
	}
	fmt.Println("---Response Cookies---")
	// TODO: How to Print Response Cookies
	for _, responseCookie := range response.Cookies() {
		fmt.Println(responseCookie.Name)
	}

	fmt.Println("---Response Body---")
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}
