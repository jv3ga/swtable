package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"backend/handlers"
)

// Middleware para manejar CORS
func enableCORS(allowedOrigin string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Permitir el origen solo si coincide con el permitido
			if origin != "" && origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// Manejar preflight requests
			if r.Method == "OPTIONS" {
				if origin == allowedOrigin {
					w.WriteHeader(http.StatusOK)
					return
				}
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	// Leer origen permitido desde variables de entorno
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		log.Fatal("ALLOWED_ORIGIN is not set")
	}

	r := mux.NewRouter()

	// Rutas de la API
	r.HandleFunc("/api/people", handlers.GetPeople).Methods("GET")
	r.HandleFunc("/api/planets", handlers.GetPlanets).Methods("GET")

	// Envolver las rutas con el middleware CORS
	http.Handle("/", enableCORS(allowedOrigin)(r))

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
