package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/sidharth-chauhan/imdb-cli/models"
)

// FetchMovieByID calls the OMDb API with an IMDb ID and returns movie details
func FetchMovieByID(imdbID string) (models.MovieDetail, error) {
	baseURL := "https://www.omdbapi.com/?apikey=YOUR_KEY&i=IMDB_ID"
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return models.MovieDetail{}, err
	}

	// Add query parameters
	q := parsedURL.Query()
	q.Set("apikey", "cad09794") // 🔐 Your actual API key
	q.Set("i", imdbID)          // 🎯 IMDb ID search
	parsedURL.RawQuery = q.Encode()

	// Make GET request
	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return models.MovieDetail{}, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.MovieDetail{}, err
	}

	// Unmarshal JSON
	var movie models.MovieDetail
	err = json.Unmarshal(body, &movie)
	if err != nil {
		return movie, err
	}

	return movie, nil
}
