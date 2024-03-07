package api

import (
	"encoding/json"
	"net/http"
	"rest-web-service/internal/clients"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// StatusHandler handles requests to the status endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := clients.FetchServiceStatus()
	respondWithJSON(w, http.StatusOK, status)
}
