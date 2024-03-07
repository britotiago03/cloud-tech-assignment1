package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-web-service/internal/model"
)

func FetchBooksByLanguage(language string) (*model.GutendexResponse2, error) {
	url := fmt.Sprintf("http://129.241.150.113:8000/books/?languages=%s", language)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request to Gutendex API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status from Gutendex API: %s", resp.Status)
	}

	var response model.GutendexResponse2
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON response: %w", err)
	}

	return &response, nil
}

func FetchBooksAndAuthorsByLanguage(language string) (int, int, error) {
	response, err := FetchBooksByLanguage(language) // Assuming this function is already defined
	if err != nil {
		return 0, 0, err
	}

	bookCount := len(response.Results)
	authorSet := make(map[string]struct{}) // Use a set to store unique authors

	for _, book := range response.Results {
		for _, author := range book.Authors {
			if author.Name != "" {
				authorSet[author.Name] = struct{}{}
			}
		}
	}

	return bookCount, len(authorSet), nil // Returns count of books and unique authors
}
