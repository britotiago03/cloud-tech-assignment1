package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-web-service/internal/model"
	"strings"
)

// FetchBookCounts takes a slice of language codes and returns book count and author count for each language
func FetchBookCounts(languages []string) ([]map[string]interface{}, error) {
	var response []map[string]interface{}
	totalBooks, err := fetchTotalBooksCount() // Fetch total books once, outside the loop
	if err != nil {
		return nil, err
	}

	for _, lang := range languages {
		url := fmt.Sprintf("http://129.241.150.113:8000/books/?languages=%s", lang)
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error making request to Gutendex API: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("received non-OK HTTP status from Gutendex API: %s", resp.Status)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}

		var books model.GutendexResponse
		if err := json.Unmarshal(body, &books); err != nil {
			return nil, fmt.Errorf("error unmarshaling JSON response: %w", err)
		}

		// Count unique authors
		authorSet := make(map[string]struct{})
		for _, book := range books.Results {
			for _, author := range book.Authors {
				// Normalize author name by trimming spaces and converting to lower case
				authorName := strings.TrimSpace(strings.ToLower(author.Name))
				if authorName != "" {
					authorSet[authorName] = struct{}{}
				}
			}
		}

		response = append(response, map[string]interface{}{
			"language": lang,
			"books":    books.Count, // Use the count from the API response
			"authors":  len(authorSet),
			"fraction": float64(books.Count) / float64(totalBooks),
		})
	}

	return response, nil
}

// fetchTotalBooksCount fetches the total number of books available in the Gutendex API
func fetchTotalBooksCount() (int, error) {
	url := "http://129.241.150.113:8000/books/"
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error making request to Gutendex API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("received non-OK HTTP status from Gutendex API: %s", resp.Status)
	}

	var result struct {
		Count int `json:"count"`
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("error unmarshaling JSON response: %w", err)
	}

	return result.Count, nil
}
