/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/sidharth-chauhan/imdb-cli/utils"
	"github.com/spf13/cobra"
)

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		id1, _ := cmd.Flags().GetString("id1")
		id2, _ := cmd.Flags().GetString("id2")
		if id1 == "" || id2 == "" {
			cmd.Println("❌ Please provide two movie IDs using --id1 and --id2")
			return
		}
		movie1, err1 := utils.FetchMovieByID(id1)
		if err1 != nil {
			cmd.Println("❌ Error fetching movie 1:", err1)
			return
		}
		movie2, err2 := utils.FetchMovieByID(id2)
		if err2 != nil {
			cmd.Println("❌ Error fetching movie 2:", err2)
			return
		}
		imdbid1, err := strconv.ParseFloat(movie1.IMDBRating, 64)
		if err != nil {
			cmd.Println("❌ Error converting movie 1 IMDBID to integer:", err)
			return
		}
		imdbid2, err := strconv.ParseFloat(movie2.IMDBRating, 64)
		if err != nil {
			cmd.Println("❌ Error converting movie 2 IMDBID to integer:", err)
			return
		}
		fmt.Printf("\n🎬 Movie 1:\n📌 %s (%s)\n", movie1.Title, movie1.IMDBID)
		fmt.Printf("⭐ Rating: %.1f\n", imdbid1)

		fmt.Printf("\n🎬 Movie 2:\n📌 %s (%s)\n", movie2.Title, movie2.IMDBID)
		fmt.Printf("⭐ Rating: %.1f\n\n", imdbid2)

		//comparing the two movies
		fmt.Println("🏁 Comparing ratings... ")

		if imdbid1 > imdbid2 {
			fmt.Printf("🏆 %s is rated higher (%.1f vs %.1f)\n", movie1.Title, imdbid1, imdbid2)
		} else if imdbid2 > imdbid1 {
			fmt.Printf("🏆 %s is rated higher (%.1f vs %.1f)\n", movie2.Title, imdbid2, imdbid1)
		} else {
			fmt.Printf("🤝 It's a tie! Both rated %.1f\n", imdbid1)
		}

	},
}

func init() {
	rootCmd.AddCommand(compareCmd)

	compareCmd.Flags().StringP("id1", "i", "", "First movie ID to compare")
	compareCmd.Flags().StringP("id2", "j", "", "Second movie ID to compare")
}
