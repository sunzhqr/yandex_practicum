package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func UserViewHandler(users map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("ID")
		if userId == "" {
			http.Error(w, "user_id is empty", http.StatusBadRequest)
			return
		}
		user, ok := users[userId]
		if !ok {
			http.Error(w, fmt.Sprintf("User not found, your user id is %s", userId), http.StatusNotFound)
			return
		}
		userJson, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "can't provide a json. Internal error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userJson)
	}
}

func StatusHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("MyHeader:", r.Header.Values("MyHeader"))
	for _, cookie := range r.Cookies() {
		fmt.Println(cookie.Name, ":", cookie.Value)
	}
	// cookieFromServer :=
	// Check if it really sets a cookie to response
	http.SetCookie(rw, &http.Cookie{
		Name:  "captcha test",
		Value: "you're human",
	})
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "ok"}`))
}

func URLHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(rw, r)
		return
	}
	r.ParseForm()
	fmt.Println("---Query Parameters---")
	for key, values := range r.URL.Query() {
		fmt.Printf("%s: %v\r\n", key, values)
	}
	fmt.Println("---Request Header---")
	for headerName, headerValues := range r.Header {
		fmt.Printf("%s: %v\r\n", headerName, headerValues)
	}
	fmt.Println("---Parsed Form Parameters---")
	for formName, formValues := range r.Form {
		fmt.Printf("%s: %v\r\n", formName, formValues)
	}
	io.WriteString(rw, "www.shorturl.com")
}
