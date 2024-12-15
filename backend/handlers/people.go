package handlers

import (
	"encoding/json"
	"net/http"
	"backend/utils"
	"log"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	sortBy := r.URL.Query().Get("sortBy")
	order := r.URL.Query().Get("order")

	peopleResponse, err := utils.FetchPeople(query, page, sortBy, order)
	if err != nil {
		log.Printf("Error fetching people: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Extraer el campo "results" como una lista
	results, ok := peopleResponse["results"].([]interface{})
	if !ok {
		http.Error(w, "invalid response format from SWAPI", http.StatusInternalServerError)
		return
	}

	// Convertir []interface{} a []map[string]interface{}
	people := make([]map[string]interface{}, len(results))
	for i, item := range results {
		people[i] = item.(map[string]interface{})
	}

	// Ordenar datos si es necesario
	sortedPeople, err := utils.SortData(people, sortBy, order)
	if err != nil {
		log.Printf("Error sorting data: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Actualizar el campo "results" con los datos ordenados
	peopleResponse["results"] = sortedPeople

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peopleResponse)
}