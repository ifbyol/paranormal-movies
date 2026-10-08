package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetCatalog returns the full movie catalog fetched from the catalog
// service, with every movie's price overridden by getCatalog().
func GetCatalog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	catalog, err := getCatalog(ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catalog)
}
