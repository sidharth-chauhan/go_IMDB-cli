/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/sidharth-chauhan/imdb-cli/db"
	"github.com/sidharth-chauhan/imdb-cli/models"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		db := db.GetDB()
		//slice to store favorite movies
		var favorites []models.Favorite
		result := db.Find(&favorites)
		if result.Error != nil {
			cmd.Println("❌ Error fetching favorite movies:", result.Error)
			return
		}
		fmt.Println("⭐ Your Favorite Movies:")
		for _, fav := range favorites {
			fmt.Printf("\n🎬 Title: %s\n📅 Year: %s\n🆔 IMDb ID: %s\n🖼️ Poster: %s\n",
				fav.Title, fav.Year, fav.IMDBID, fav.Poster)
		}
		fmt.Println("\nTotal Favorites:", len(favorites))

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
