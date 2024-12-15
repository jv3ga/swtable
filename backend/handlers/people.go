package handlers

import (
	"backend/utils"
	"net/http"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	// Obtener parámetros de la URL
	query := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	sortBy := r.URL.Query().Get("sortBy")
	order := r.URL.Query().Get("order")

	// Llamar a la utilidad para obtener los datos
	utils.FetchPeople(w, query, page, sortBy, order)
}
