package test

import (
	"testing"

	"github.com/sidharth-chauhan/imdb-cli/models"
	"github.com/sidharth-chauhan/imdb-cli/utils"
)

func TestFetchMovieByID(t *testing.T) {
	// Use a known valid IMDb ID for testing
	imdbID := "tt0111161" // The Shawshank Redemption
	movie, err := utils.FetchMovieByID(imdbID)
	if err != nil {
		t.Fatalf("FetchMovieByID failed: %v", err)
	}
	if movie.IMDBID != imdbID {
		t.Errorf("Expected IMDBID %s, got %s", imdbID, movie.IMDBID)
	}
	if movie.Title == "" {
		t.Error("Expected non-empty Title")
	}
}

func TestFavoriteModel(t *testing.T) {
	fav := models.Favorite{
		Title:  "Test Movie",
		Year:   "2024",
		IMDBID: "tt1234567",
		Type:   "movie",
		Poster: "http://example.com/poster.jpg",
	}
	if fav.Title != "Test Movie" {
		t.Error("Favorite Title mismatch")
	}
}
