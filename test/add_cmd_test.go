package test

import (
	"testing"

	"github.com/sidharth-chauhan/imdb-cli/db"
	"github.com/sidharth-chauhan/imdb-cli/models"
	"github.com/sidharth-chauhan/imdb-cli/utils"
)

// clearFavoritesTable is removed and replaced with shared ClearFavoritesTable.

func TestAddFavoriteValid(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	d := db.GetDB()
	imdbID := "tt0111161"
	d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
	movie, err := utils.FetchMovieByID(imdbID)
	if err != nil || movie.Response != "True" {
		t.Fatalf("FetchMovieByID failed: %v", err)
	}
	fav := models.Favorite{
		Title:  movie.Title,
		Year:   movie.Year,
		IMDBID: movie.IMDBID,
		Type:   movie.Type,
		Poster: movie.Poster,
	}
	if err := d.Create(&fav).Error; err != nil {
		t.Fatalf("Failed to create favorite: %v", err)
	}
	// Check that favorite is in DB
	var found models.Favorite
	if err := d.Where("imdb_id = ?", imdbID).First(&found).Error; err != nil {
		t.Error("Favorite not found after add")
	}
	// Clean up
	d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
}

func TestAddFavoriteDuplicate(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	d := db.GetDB()
	imdbID := "tt0111161"
	d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
	movie, _ := utils.FetchMovieByID(imdbID)
	fav := models.Favorite{
		Title:  movie.Title,
		Year:   movie.Year,
		IMDBID: movie.IMDBID,
		Type:   movie.Type,
		Poster: movie.Poster,
	}
	_ = d.Create(&fav).Error
	dup := models.Favorite{
		Title:  movie.Title,
		Year:   movie.Year,
		IMDBID: movie.IMDBID,
		Type:   movie.Type,
		Poster: movie.Poster,
	}
	err := d.Create(&dup).Error
	if err == nil {
		t.Error("Expected error when adding duplicate favorite, got nil")
	}
	// Ensure only one entry exists
	var count int64
	d.Model(&models.Favorite{}).Where("imdb_id = ?", imdbID).Count(&count)
	if count != 1 {
		t.Errorf("Expected 1 favorite, got %d", count)
	}
	d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
}

func TestAddFavoriteInvalidID(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	invalidID := "tt0000000"
	movie, err := utils.FetchMovieByID(invalidID)
	if err != nil {
		// Acceptable: network or JSON error, skip test
		t.Skip("Skipping: network or JSON error for invalid IMDb ID")
		return
	}
	if movie.Response != "False" || movie.Error == "" {
		t.Errorf("Expected Response=False and non-empty Error for invalid IMDb ID, got Response=%s, Error=%s", movie.Response, movie.Error)
	}
}
