package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

const BaseURL = "https://swapi.py4e.com/api/"

// APIResponse defines the standard structure for API responses
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// FetchFromSWAPI retrieves data from SWAPI with filters
func FetchFromSWAPI(w http.ResponseWriter, resource, query, page, sortBy, order string) {
	url := fmt.Sprintf("%s/%s/?search=%s&page=%s", BaseURL, resource, query, page)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error making request to SWAPI: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("SWAPI returned non-200 status: %d", resp.StatusCode), resp.StatusCode)
		return
	}

	// Decodificar la respuesta JSON
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding SWAPI response: %v", err), http.StatusInternalServerError)
		return
	}

	// Validar y procesar el campo "results"
	results, ok := data["results"].([]interface{})
	if !ok {
		http.Error(w, "Invalid response format from SWAPI: missing or invalid 'results' field", http.StatusInternalServerError)
		return
	}

	// Convertir []interface{} a []map[string]interface{}
	items := make([]map[string]interface{}, len(results))
	for i, item := range results {
		itemMap, valid := item.(map[string]interface{})
		if !valid {
			http.Error(w, "Invalid item format in results", http.StatusInternalServerError)
			return
		}
		items[i] = itemMap
	}

	// Aplicar ordenación si se especifica
	if sortBy != "" {
		sortedItems, err := SortData(items, sortBy, order)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error sorting data: %v", err), http.StatusBadRequest)
			return
		}
		data["results"] = sortedItems
	} else {
		data["results"] = items
	}

	// Enviar la respuesta ordenada
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func FetchPeople(w http.ResponseWriter, query, page, sortBy, order string) {
	FetchFromSWAPI(w, "people", query, page, sortBy, order)
}

func FetchPlanets(w http.ResponseWriter, query, page, sortBy, order string) {
	FetchFromSWAPI(w, "planets", query, page, sortBy, order)
}

func SortData(data []map[string]interface{}, sortBy, order string) ([]map[string]interface{}, error) {
	// Validar la dirección del orden (ascendente o descendente)
	ascending := order != "desc"

	// Realizar la ordenación
	sort.Slice(data, func(i, j int) bool {
		// Comparar los valores de las claves `sortBy`
		valI, okI := data[i][sortBy]
		valJ, okJ := data[j][sortBy]
		if !okI || !okJ {
			return ascending
		}

		// Manejar diferentes tipos de valores
		switch valI := valI.(type) {
		case string:
			valJ, _ := valJ.(string)
			if ascending {
				return valI < valJ
			}
			return valI > valJ
		case float64:
			valJ, _ := valJ.(float64)
			if ascending {
				return valI < valJ
			}
			return valI > valJ
		default:
			return ascending
		}
	})

	return data, nil
}
