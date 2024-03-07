package api

import (
	"encoding/json"
	"net/http"
	"rest-web-service/internal/clients"
	"rest-web-service/internal/model"
	"strconv"
	"strings"
)

func ReadershipHandler(w http.ResponseWriter, r *http.Request) {
	// Split the URL path into segments
	pathSegments := strings.Split(r.URL.Path, "/")
	// Ensure there are enough segments to include the language parameter
	if len(pathSegments) < 3 {
		http.Error(w, "URL path must include a language parameter", http.StatusBadRequest)
		return
	}
	// Assuming the language parameter is the last segment in the URL path
	language := pathSegments[len(pathSegments)-1]

	// Extract the limit parameter from the query string, if present
	limitStr := r.URL.Query().Get("limit")
	var limit int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	bookCount, authorCount, err := clients.FetchBooksAndAuthorsByLanguage(language)
	if err != nil {
		http.Error(w, "Failed to fetch books and authors: "+err.Error(), http.StatusInternalServerError)
		return
	}

	countries, err := clients.FetchCountriesByLanguage(language)
	if err != nil {
		http.Error(w, "Failed to fetch countries: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Apply the limit to the number of countries, if specified
	if limit > 0 && limit < len(countries) {
		countries = countries[:limit]
	}

	var readershipData []model.ReadershipResponse
	for _, country := range countries {
		population, err := clients.FetchPopulationByCountryCode(country.IsoCode)
		if err != nil {
			http.Error(w, "Failed to fetch population: "+err.Error(), http.StatusInternalServerError)
			return
		}

		readershipData = append(readershipData, model.ReadershipResponse{
			Country:    country.OfficialName,
			IsoCode:    country.IsoCode,
			Books:      bookCount,
			Authors:    authorCount,
			Readership: population,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(readershipData); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
