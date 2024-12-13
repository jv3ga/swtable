package utils

import (
	"testing"
)

func TestSortData(t *testing.T) {
	data := []map[string]interface{}{
		{"name": "Luke", "created": "2024-01-01"},
		{"name": "Leia", "created": "2024-01-03"},
		{"name": "Han", "created": "2024-01-02"},
	}

	// Ordenar por nombre ascendente
	sortedData, err := SortData(data, "name", "asc")
	if err != nil {
		t.Fatal(err)
	}

	if sortedData[0]["name"] != "Han" {
		t.Errorf("expected Han, got %v", sortedData[0]["name"])
	}

	// Ordenar por fecha descendente
	sortedData, err = SortData(data, "created", "desc")
	if err != nil {
		t.Fatal(err)
	}

	if sortedData[0]["created"] != "2024-01-03" {
		t.Errorf("expected 2024-01-03, got %v", sortedData[0]["created"])
	}
}
