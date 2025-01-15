package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	chiRouter := chi.NewRouter()
	chiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("chi"))
	})
	chiRouter.Get("/item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		io.WriteString(w, fmt.Sprintf("item = %s", id))
	})
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", chiRouter)
}
