package main

import (
	"backend/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Middleware para manejar CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Rutas de la API
	r.HandleFunc("/api/people", handlers.GetPeople).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/planets", handlers.GetPlanets).Methods("GET", "OPTIONS")

	// Envolver las rutas con el middleware CORS
	http.Handle("/", enableCORS(r))

	log.Println("My Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
