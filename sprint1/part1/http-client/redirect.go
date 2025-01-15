package main

import (
	"io"
	"net/http"
)

func mainRedirect() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			pl(req.URL)
			return nil
		},
	}
	response, err := client.Get("http://ya.ru")
	if err != nil {
		pl(err)
		return
	}
	_, err = io.Copy(io.Discard, response.Body)
	response.Body.Close()
	if err != nil {
		pl(err)
		return
	}
	pl(response.StatusCode, "OK")
}
