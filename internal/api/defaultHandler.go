package api

import (
	"fmt"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	output := "Invalid endpoint, please use one of the following:\n" +
		"/librarystats/v1/bookcount/\n" +
		"/librarystats/v1/readership/{language}\n" +
		"/librarystats/v1/status/"

	_, err := fmt.Fprint(w, output) // Corrected: removed format specifier and passed output directly

	if err != nil {
		http.Error(w, "Error occurred", http.StatusInternalServerError)
	}
}
