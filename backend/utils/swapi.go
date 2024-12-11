package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// Sorting logic here (optional for demonstration)

	return data, nil
}
