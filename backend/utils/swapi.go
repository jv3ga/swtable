package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
	"sort"
	"strings"
)

const BaseURL = "https://swapi.dev/api"

// FetchPeople retrieves people from SWAPI with filters
func FetchPeople(query, page, sortBy, order string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/people/?search=%s&page=%s", BaseURL, query, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// Sorting logic here (optional for demonstration)

	return data, nil
}


// FetchPlanets retrieves planets from SWAPI with filters
func FetchPlanets(query, page, sortBy, order string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/planets/?search=%s&page=%s", BaseURL, query, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
// SortData ordena una lista de mapas por un campo dado y un orden (asc o desc).
func SortData(data []map[string]interface{}, sortBy string, order string) ([]map[string]interface{}, error) {
	if sortBy == "" {
		return data, nil // Si no se especifica sortBy, devuelve los datos tal cual.
	}

	// Convertir el orden a minúsculas para consistencia
	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		return nil, errors.New("invalid order: must be 'asc' or 'desc'")
	}

	// Función de comparación genérica
	sort.Slice(data, func(i, j int) bool {
		// Extraer los valores de los campos
		val1, ok1 := data[i][sortBy]
		val2, ok2 := data[j][sortBy]

		// Si el campo no existe, considera que no hay orden
		if !ok1 || !ok2 {
			return false
		}

		// Comparar como cadenas
		str1, str1Ok := val1.(string)
		str2, str2Ok := val2.(string)

		if str1Ok && str2Ok {
			if order == "asc" {
				return str1 < str2
			}
			return str1 > str2
		}

		// Comparar como fechas
		if order == "asc" {
			return val1.(string) < val2.(string)
		}
		return val1.(string) > val2.(string)
	})

	return data, nil
}