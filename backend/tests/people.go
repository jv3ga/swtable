package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"backend/handlers"
)

func TestGetPeople(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/people?search=luke&page=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPeople)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if len(rr.Body.String()) == 0 {
		t.Errorf("handler returned empty body")
	}
}
