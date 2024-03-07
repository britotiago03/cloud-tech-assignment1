package clients

import (
	"fmt"
	"net/http"
	"rest-web-service/internal/model"
	"time"
)

// Assume startTime is a global variable that stores when the service was started
var startTime = time.Now()

func checkServiceStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Service Unreachable"
	}
	defer resp.Body.Close()

	// Return the status code as a string
	return fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}

// FetchServiceStatus checks the status of dependent services and returns their status
func FetchServiceStatus() model.ServiceStatus {
	// URLs of the dependent services
	gutendexAPI := "http://129.241.150.113:8000/"
	languageAPI := "http://129.241.150.113:3000/"
	countriesAPI := "http://129.241.150.113:8080/v3.1/"

	// Check the status of each service
	gutendexStatus := checkServiceStatus(gutendexAPI)
	languageStatus := checkServiceStatus(languageAPI)
	countriesStatus := checkServiceStatus(countriesAPI)

	// Calculate uptime
	uptime := time.Since(startTime).String()

	// Construct the response
	status := model.ServiceStatus{
		GutendexAPIStatus:  gutendexStatus,
		LanguageAPIStatus:  languageStatus,
		CountriesAPIStatus: countriesStatus,
		Version:            "v1", // Your service version
		Uptime:             uptime,
	}

	return status
}
