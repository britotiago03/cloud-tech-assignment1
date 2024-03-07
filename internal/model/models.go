package model

// Book represents the basic structure of a book as returned by the Gutendex API.
type Book struct {
	Title   string   `json:"title"`
	Authors []Author `json:"authors"`
}

// Author represents the structure of an author associated with a book in the Gutendex API response.
type Author struct {
	Name string `json:"name"`
}

// GutendexResponse represents the response structure from the Gutendex API for a list of books.
type GutendexResponse struct {
	Count   int `json:"count"`
	Results []struct {
		Authors []struct {
			Name string `json:"name"`
		} `json:"authors"`
	} `json:"results"`
}

type GutendexResponse2 struct {
	Results []Book `json:"results"`
}

// BookCountResponse represents the structure of the response for the book count endpoint of your API.
// It includes the language code, the count of books and authors for that language, and the fraction of books in that language compared to the total number of books.
type BookCountResponse struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fraction"`
}

// CountryInfo represents the structure of a country as it relates to its language and other attributes in your application.
// This structure is used when fetching countries that speak a given language from an external API.
type CountryInfo struct {
	IsoCode       string `json:"ISO3166_1_Alpha_2"`
	OfficialName  string `json:"Official_Name"`
	RegionName    string `json:"Region_Name,omitempty"` // omitempty is used to allow this field to be absent in the JSON without causing errors
	SubRegionName string `json:"Sub_Region_Name,omitempty"`
	Language      string `json:"Language,omitempty"`
}

// Country represents the basic structure of a country as used in your application, typically fetched from an external API.
type Country struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Population int `json:"population"`
}

// ReadershipResponse represents the structure for the readership data response.
// It includes country information, the count of books and authors for a given language, and the readership potential based on the country's population.
type ReadershipResponse struct {
	Country    string `json:"country"`    // The name of the country
	IsoCode    string `json:"isocode"`    // The ISO code of the country
	Books      int    `json:"books"`      // The count of books available in the language
	Authors    int    `json:"authors"`    // The count of unique authors for the books
	Readership int    `json:"readership"` // The potential readership based on the country's population
}

// ServiceStatus represents the health status of various external services your API depends on.
// It includes the HTTP status of the Gutendex API, language-to-countries mapping API, and any other relevant services.
type ServiceStatus struct {
	GutendexAPIStatus  string `json:"gutendexapi"`  // The status of the Gutendex API
	LanguageAPIStatus  string `json:"languageapi"`  // The status of the language-to-countries mapping API
	CountriesAPIStatus string `json:"countriesapi"` // The status of the countries API
	Version            string `json:"version"`      // The version of your service
	Uptime             string `json:"uptime"`       // The uptime of your service since the last restart
}
