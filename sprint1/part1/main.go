package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Subject struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(".."))
	fmt.Println(fs)
	mux.Handle(`/repo`, http.StripPrefix("/repo", fs))
	mux.HandleFunc("/thing", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../responseFolder/res.go")
	})
	mux.HandleFunc("/request", requestPage)
	mux.HandleFunc("/json", JsonHandler)
	mux.Handle("/redir", http.RedirectHandler("https://about.gitlab.com/", http.StatusMovedPermanently))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func requestPage(w http.ResponseWriter, r *http.Request) {
	body := fmt.Sprintf("Request Method: %s\n", r.Method)
	body += "-----Header-----"
	for k, v := range r.Header {
		body += fmt.Sprintf("%s: %v\n", k, v)
	}
	body += "-----Query Parameters-----\n"
	// for k, v := range r.URL.Query() {
	// 	body += fmt.Sprintf("%s: %v\n", k, v)
	// }
	if err := r.ParseForm(); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	for formKey, FormValue := range r.Form {
		body += fmt.Sprintf("%s: %v\n", formKey, FormValue)
	}
	w.Write([]byte(body))
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	subj := Subject{"Course", 50000}
	res, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
