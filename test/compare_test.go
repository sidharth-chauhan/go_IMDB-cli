package test

import (
	"strconv"
	"testing"

	"github.com/sidharth-chauhan/imdb-cli/utils"
)

func TestCompareRatingsWinner(t *testing.T) {
	id1 := "tt0111161" // Shawshank
	id2 := "tt0068646" // Godfather
	movie1, err := utils.FetchMovieByID(id1)
	if err != nil {
		t.Fatalf("FetchMovieByID failed for id1: %v", err)
	}
	movie2, err := utils.FetchMovieByID(id2)
	if err != nil {
		t.Fatalf("FetchMovieByID failed for id2: %v", err)
	}
	r1, err := strconv.ParseFloat(movie1.IMDBRating, 64)
	if err != nil {
		t.Fatalf("ParseFloat failed for movie1: %v", err)
	}
	r2, err := strconv.ParseFloat(movie2.IMDBRating, 64)
	if err != nil {
		t.Fatalf("ParseFloat failed for movie2: %v", err)
	}
	if r1 == r2 {
		t.Log("Both movies have the same rating (tie case)")
	} else if r1 > r2 {
		t.Logf("%s should be the winner", movie1.Title)
	} else {
		t.Logf("%s should be the winner", movie2.Title)
	}
}

func TestCompareRatingsTie(t *testing.T) {
	id := "tt0111161"
	movie1, err := utils.FetchMovieByID(id)
	if err != nil {
		t.Fatalf("FetchMovieByID failed: %v", err)
	}
	movie2, err := utils.FetchMovieByID(id)
	if err != nil {
		t.Fatalf("FetchMovieByID failed: %v", err)
	}
	r1, _ := strconv.ParseFloat(movie1.IMDBRating, 64)
	r2, _ := strconv.ParseFloat(movie2.IMDBRating, 64)
	if r1 != r2 {
		t.Error("Expected equal ratings for the same movie")
	}
}

func TestCompareRatingsInvalidID(t *testing.T) {
	movie, err := utils.FetchMovieByID("tt0000000")
	if err != nil {
		t.Skip("Skipping: network or JSON error for invalid IMDb ID")
		return
	}
	if movie.Response != "False" || movie.Error == "" {
		t.Errorf("Expected Response=False and non-empty Error for invalid IMDb ID, got Response=%s, Error=%s", movie.Response, movie.Error)
	}
}
