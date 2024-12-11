package handlers

import (
	"encoding/json"
	"net/http"
	"backend/utils"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	sortBy := r.URL.Query().Get("sortBy")
	order := r.URL.Query().Get("order")

	people, err := utils.FetchPeople(query, page, sortBy, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
