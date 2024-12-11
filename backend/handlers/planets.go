package handlers

import (
	"encoding/json"
	"net/http"
	"backend/utils"
)

// GetPlanets handles the planets API endpoint
func GetPlanets(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	sortBy := r.URL.Query().Get("sortBy")
	order := r.URL.Query().Get("order")

	planets, err := utils.FetchPlanets(query, page, sortBy, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(planets)
}
