package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-web-service/internal/model"
)

func FetchCountriesByLanguage(language string) ([]model.CountryInfo, error) {
	url := fmt.Sprintf("http://129.241.150.113:3000/language2countries/%s", language)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching countries: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	var countries []model.CountryInfo
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return countries, nil
}
