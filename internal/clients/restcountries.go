package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchPopulationByCountryCode(isoCode string) (int, error) {
	url := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", isoCode)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error fetching population: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	var result []struct {
		Population int `json:"population"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if len(result) == 0 {
		return 0, fmt.Errorf("no data returned for ISO code: %s", isoCode)
	}

	return result[0].Population, nil
}
