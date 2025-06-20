package test

import (
	"testing"

	"github.com/sidharth-chauhan/imdb-cli/db"
	"github.com/sidharth-chauhan/imdb-cli/models"
)

func TestInitDB(t *testing.T) {
	db.InitDB()
	d := db.GetDB()
	if d == nil {
		t.Fatal("DB connection is nil after InitDB")
	}
}

func TestFavoriteCRUD(t *testing.T) {
	db.InitDB()
	d := db.GetDB()
	fav := models.Favorite{
		Title:  "Test Movie",
		Year:   "2024",
		IMDBID: "tt7654321",
		Type:   "movie",
		Poster: "http://example.com/poster.jpg",
	}
	// Create
	if err := d.Create(&fav).Error; err != nil {
		t.Fatalf("Failed to create favorite: %v", err)
	}
	// Read
	var found models.Favorite
	if err := d.Where("imdb_id = ?", fav.IMDBID).First(&found).Error; err != nil {
		t.Fatalf("Failed to find favorite: %v", err)
	}
	// Delete
	if err := d.Where("imdb_id = ?", fav.IMDBID).Delete(&models.Favorite{}).Error; err != nil {
		t.Fatalf("Failed to delete favorite: %v", err)
	}
}
