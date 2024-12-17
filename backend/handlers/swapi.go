package handlers

import (
	"backend/utils"
	"net/http"
)

// APIError estructura para mensajes de error
type APIError struct {
	Message string `json:"message"`
}

// getQueryStringValues extrae valores de la URL
func getQueryStringValues(r *http.Request) (string, string, string, string) {
	// Obtener parámetros de la URL
	query := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	sortBy := r.URL.Query().Get("sortBy")
	order := r.URL.Query().Get("order")
	return query, page, sortBy, order
}

// GetPlanets maneja la ruta /api/planets
func GetPlanets(w http.ResponseWriter, r *http.Request) {
	query, page, sortBy, order := getQueryStringValues(r)
	utils.FetchFromSWAPI(w, "planets", query, page, sortBy, order)
}

// GetPeople maneja la ruta /api/people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	query, page, sortBy, order := getQueryStringValues(r)
	utils.FetchFromSWAPI(w, "people", query, page, sortBy, order)
}
