package test

import "github.com/sidharth-chauhan/imdb-cli/db"

// ClearFavoritesTable removes all rows from the favorites table for test isolation.
func ClearFavoritesTable() {
	d := db.GetDB()
	d.Exec("DELETE FROM favorites")
}
