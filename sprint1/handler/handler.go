package main

import (
	"encoding/json"
	"fmt"
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
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "ok"}`))
}
