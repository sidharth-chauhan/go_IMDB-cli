package test

import (
	"testing"

	"github.com/sidharth-chauhan/imdb-cli/db"
	"github.com/sidharth-chauhan/imdb-cli/models"
)

func TestDeleteFavoriteValid(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	d := db.GetDB()
	imdbID := "tt9999999"
	fav := models.Favorite{
		Title:  "Test Movie",
		Year:   "2024",
		IMDBID: imdbID,
		Type:   "movie",
		Poster: "",
	}
	if err := d.Create(&fav).Error; err != nil {
		t.Fatalf("Failed to create favorite: %v", err)
	}
	result := d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
	if result.Error != nil {
		t.Fatalf("Failed to delete favorite: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		t.Error("Expected RowsAffected > 0 when deleting existing favorite")
	}
	// Ensure it's gone
	var found models.Favorite
	err := d.Where("imdb_id = ?", imdbID).First(&found).Error
	if err == nil {
		t.Error("Favorite still found after delete")
	}
}

func TestDeleteFavoriteNonExistent(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	d := db.GetDB()
	imdbID := "tt0000000"
	result := d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
	if result.Error != nil {
		t.Fatalf("Unexpected error deleting non-existent favorite: %v", result.Error)
	}
	if result.RowsAffected != 0 {
		t.Error("Expected RowsAffected == 0 when deleting non-existent favorite")
	}
}

func TestDeleteFavoriteTwice(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	d := db.GetDB()
	imdbID := "tt7777777"
	fav := models.Favorite{
		Title:  "Test Movie Twice",
		Year:   "2024",
		IMDBID: imdbID,
		Type:   "movie",
		Poster: "",
	}
	if err := d.Create(&fav).Error; err != nil {
		t.Fatalf("Failed to create favorite: %v", err)
	}
	_ = d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
	result := d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
	if result.RowsAffected != 0 {
		t.Error("Expected RowsAffected == 0 when deleting already deleted favorite")
	}
}
