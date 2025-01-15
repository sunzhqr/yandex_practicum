package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func mainOver() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Response Status Code: %d\r\n", response.StatusCode)
	fmt.Println("-----HEADER-----")
	// for k, v := range response.Header {
	// 	fmt.Printf("Header: %s, Value: %v\r\n", k, v[0])
	// }
	contentTypeValue := response.Header.Get("Content-Type")
	contentTypeValues := response.Header.Values("Content-Type")
	fmt.Printf("Header: Content-Type, Value: %v\r\n", contentTypeValue)
	fmt.Printf("Header: Content-Type, Values: %v\r\n", contentTypeValues)
	fmt.Println("-----RESPONSE BODY-----")
	// respBody, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if len(respBody) > 512 {
	// 	respBody = respBody[:512]
	// }
	// fmt.Println("Response:", string(respBody))
	// fmt.Println("Response Length:", len(string(respBody)))

	defer response.Body.Close()
	var written int64
	if written, err = io.CopyN(os.Stdout, response.Body, 512); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\nResponse Length:", written)
}
