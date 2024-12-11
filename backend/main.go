package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"backend/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/people", handlers.GetPeople).Methods("GET")
	r.HandleFunc("/api/planets", handlers.GetPlanets).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
