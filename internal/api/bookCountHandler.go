package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"rest-web-service/internal/clients"
	"strings"
)

func validLanguage(language string) bool {
	valid, _ := regexp.MatchString("^[a-zA-Z]{2}$", language)
	return valid
}

// BookCountHandler handles requests to the bookcount endpoint
func BookCountHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	languagesParam, ok := query["language"]
	if !ok || len(languagesParam[0]) < 1 {
		http.Error(w, "Language parameter is missing", http.StatusBadRequest)
		return
	}
	languages := strings.Split(languagesParam[0], ",")

	for _, language := range languages {
		if !validLanguage(language) {
			err := fmt.Sprintf("An invalid language was provided. " +
				"Please check that all languages specified are " +
				"part of ISO 639 Set 1")
			http.Error(w, err, http.StatusBadRequest)
			return
		}
	}

	bookCounts, err := clients.FetchBookCounts(languages)
	if err != nil {
		http.Error(w, "Failed to fetch book counts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookCounts); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
