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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		imdbid, _ := cmd.Flags().GetString("id")
		if imdbid == "" {
			cmd.Println("❌ Please provide a movie ID using --id or -i")
			return
		}
		db := db.GetDB()

		result := db.Where("imdb_id = ?", imdbid).Delete(&models.Favorite{})
		if result.Error != nil {
			cmd.Println("❌ Error deleting movie:", result.Error)
			return
		}
		if result.RowsAffected == 0 {
			cmd.Println("❌ Movie not found in your favorites.")
			return
		}
		fmt.Printf("✅ Movie with ID '%s' has been deleted from your favorites!\n", imdbid)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("id", "i", "", "IMDB ID of the movie to delete")
}
