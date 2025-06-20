/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/sidharth-chauhan/imdb-cli/db"
	"github.com/sidharth-chauhan/imdb-cli/models"
	"github.com/sidharth-chauhan/imdb-cli/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		imdbid, _ := cmd.Flags().GetString("id")
		if imdbid == "" {
			cmd.Println("❌ Please provide a movie ID using --id or -i")
			return
		}
		db := db.GetDB()
		movie, err := utils.FetchMovieByID(imdbid)
		if err != nil {
			cmd.Println("❌ Error fetching movie:", err)
			return
		}
		var fav models.Favorite
		result := db.Where("imdb_id = ?", movie.IMDBID).First(&fav)
		if result.Error == nil {
			cmd.Println("❌ Movie already exists in your favorites.")
			return
		}
		//new struct for favorite movie
		fav = models.Favorite{
			Title:  movie.Title,
			Year:   movie.Year,
			IMDBID: movie.IMDBID,
			Type:   movie.Type,
			Poster: movie.Poster,
		}
		err = db.Create(&fav).Error
		if err != nil {
			cmd.Println("❌ Error adding movie to favorites:", err)
			return
		}
		fmt.Printf("✅ Movie '%s' (%s) added to your favorites!\n", movie.Title, movie.IMDBID)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("id", "i", "", "IMDB ID of the movie to add")
}
