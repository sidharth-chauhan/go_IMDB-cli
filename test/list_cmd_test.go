package test

import (
	"testing"

	"github.com/sidharth-chauhan/imdb-cli/db"
	"github.com/sidharth-chauhan/imdb-cli/models"
)

func TestListFavoritesEmpty(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	d := db.GetDB()
	var favorites []models.Favorite
	result := d.Find(&favorites)
	if result.Error != nil {
		t.Fatalf("Error fetching favorites: %v", result.Error)
	}
	if len(favorites) != 0 {
		t.Errorf("Expected 0 favorites, got %d", len(favorites))
	}
}

func TestListFavoritesSingleAndDelete(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	d := db.GetDB()
	imdbID := "tt8888888"
	fav := models.Favorite{
		Title:  "Test Movie 2",
		Year:   "2024",
		IMDBID: imdbID,
		Type:   "movie",
		Poster: "",
	}
	if err := d.Create(&fav).Error; err != nil {
		t.Fatalf("Failed to create favorite: %v", err)
	}
	var favorites []models.Favorite
	result := d.Find(&favorites)
	if result.Error != nil {
		t.Fatalf("Error fetching favorites: %v", result.Error)
	}
	if len(favorites) != 1 {
		t.Errorf("Expected 1 favorite, got %d", len(favorites))
	}
	if favorites[0].IMDBID != imdbID {
		t.Errorf("Expected IMDBID %s, got %s", imdbID, favorites[0].IMDBID)
	}
	// Delete and check empty again
	d.Where("imdb_id = ?", imdbID).Delete(&models.Favorite{})
	var afterDelete []models.Favorite
	d.Find(&afterDelete)
	if len(afterDelete) != 0 {
		t.Error("Expected 0 favorites after delete")
	}
}

func TestListFavoritesMultiple(t *testing.T) {
	db.InitDB()
	ClearFavoritesTable()
	d := db.GetDB()
	imdbIDs := []string{"tt1111111", "tt2222222", "tt3333333"}
	for _, id := range imdbIDs {
		fav := models.Favorite{
			Title:  "Movie " + id,
			Year:   "2024",
			IMDBID: id,
			Type:   "movie",
			Poster: "",
		}
		if err := d.Create(&fav).Error; err != nil {
			t.Fatalf("Failed to create favorite: %v", err)
		}
	}
	var favorites []models.Favorite
	result := d.Find(&favorites)
	if result.Error != nil {
		t.Fatalf("Error fetching favorites: %v", result.Error)
	}
	if len(favorites) != len(imdbIDs) {
		t.Errorf("Expected %d favorites, got %d", len(imdbIDs), len(favorites))
	}
	// Check all IDs present
	idMap := map[string]bool{}
	for _, f := range favorites {
		idMap[f.IMDBID] = true
	}
	for _, id := range imdbIDs {
		if !idMap[id] {
			t.Errorf("Favorite with IMDBID %s not found", id)
		}
	}
	ClearFavoritesTable()
}
