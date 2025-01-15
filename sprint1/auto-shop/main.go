package main

import (
	"io"
	"log"
	"net/http"
	"strings"
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

func carsHandle(w http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	io.WriteString(w, strings.Join(carsList, ", "))
}

func carsListFunc() []string {
	var carsList []string
	for _, car := range cars {
		carsList = append(carsList, car)
	}
	return carsList
}

func carHandle(w http.ResponseWriter, r *http.Request) {
	carID := r.URL.Query().Get("id")
	if carID == "" {
		http.Error(w, "carID param is missed", http.StatusBadRequest)
		return
	}
	w.Write([]byte(carFunc(carID)))
}

func carFunc(id string) string {
	if car, ok := cars[id]; ok {
		return car
	}
	return "unknown identifier: " + id
}

func main() {
	http.HandleFunc("/cars", carsHandle)
	http.HandleFunc("/car", carHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
