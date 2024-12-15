package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

const BaseURL = "https://swapi.dev/api"

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// FetchFromSWAPI retrieves data from SWAPI with filters
func FetchFromSWAPI(resource, query, page string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/?search=%s&page=%s", BaseURL, resource, query, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, &HTTPError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error making request to SWAPI: %v", err),
		}
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("SWAPI returned non-200 status: %d", resp.StatusCode),
		}
	}

	// Decodificar la respuesta JSON
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, &HTTPError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error decoding SWAPI response: %v", err),
		}
	}

	return data, nil
}

// FetchPeople retrieves people from SWAPI with filters
func FetchPeople(query, page, sortBy, order string) (map[string]interface{}, error) {
	return FetchFromSWAPI("people", query, page)
}

// FetchPlanets retrieves planets from SWAPI with filters
func FetchPlanets(query, page, sortBy, order string) (map[string]interface{}, error) {
	return FetchFromSWAPI("people", query, page)
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
