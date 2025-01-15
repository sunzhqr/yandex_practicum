package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

func carsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is allowed", http.StatusMethodNotAllowed)
		return
	}
	carsList := carsFunc()
	io.WriteString(w, strings.Join(carsList, ", "))
}

func brandsHandler(w http.ResponseWriter, r *http.Request) {
	brand := strings.ToLower(chi.URLParam(r, "brand"))
	carsOfBrand := make([]string, 0, len(cars))
	for _, car := range cars {
		if carBrand := strings.Split(car, " ")[0]; strings.ToLower(carBrand) == brand {
			carsOfBrand = append(carsOfBrand, car)
		}
	}
	io.WriteString(w, strings.Join(carsOfBrand, ", "))
}

func modelsHandler(w http.ResponseWriter, r *http.Request) {
	model := strings.ToLower(chi.URLParam(r, "brand") + ` ` + chi.URLParam(r, "model"))
	for _, car := range cars {
		if strings.ToLower(car) == model {
			w.Write([]byte(car))
			return
		}
	}
	http.Error(w, "unknown model: "+model, http.StatusNotFound)
}

func carsFunc() []string {
	carsList := make([]string, 0, len(cars))
	for _, car := range cars {
		carsList = append(carsList, car)
	}
	return carsList
}

func carHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is allowed", http.StatusMethodNotAllowed)
		return
	}
	carID := chi.URLParam(r, "id")
	if carID == "" {
		http.Error(w, "id doesn't privided", http.StatusBadRequest)
		return
	}
	car, ok := cars[carID]
	if !ok {
		http.Error(w, "content hasn't found", http.StatusNotFound)
		return
	}
	w.Write([]byte(car))

}

func main() {
	router := CarRouter() // chi.NewRouter()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func CarRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/cars", func(router chi.Router) {
		router.Get("/", carsHandler)
		router.Route("/{brand}", func(router chi.Router) {
			router.Get("/", brandsHandler)
			router.Get("/{model}", modelsHandler)
		})
	})
	router.Get("/car/{id}", carHandler)
	return router
}
