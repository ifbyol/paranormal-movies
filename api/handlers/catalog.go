package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// DefaultCatalogPrice is the fixed price (in USD) returned for every movie
// in the catalog.
const DefaultCatalogPrice = 100.0

// GetCatalog fetches the catalog from the catalog service and overrides
// every movie's price with DefaultCatalogPrice before returning it.
func GetCatalog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	catalog, err := getCatalog(ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i := range catalog {
		catalog[i].Price = DefaultCatalogPrice
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catalog)
}
