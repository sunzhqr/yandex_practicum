package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var pl = fmt.Println

func main() {
	endpoint := "http://localhost:8080/url"
	data := url.Values{}
	pl("Input a long URL")
	reader := bufio.NewReader(os.Stdin)
	longURL, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	longURL = strings.TrimSuffix(longURL, "\n")
	data.Set("url", longURL)
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	pl("Status Code:", response.StatusCode)
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	pl(string(responseBody))
}
