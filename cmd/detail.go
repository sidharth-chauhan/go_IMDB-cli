/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

type MovieDetail struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	IMDBRating string `json:"imdbRating"`
	IMDBID     string `json:"imdbID"`
	Response   string `json:"Response"`
	Error      string `json:"Error"`
}

// detailCmd represents the detail command
var detailCmd = &cobra.Command{
	Use:   "detail",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		idimdb, _ := cmd.Flags().GetString("id")
		if idimdb == "" {
			fmt.Println("❌ Please provide a movie ID using --id or -i")
			return
		}
		movieURL := "https://www.omdbapi.com/?apikey=YOUR_KEY&i=IMDB_ID"
		//string url to url object (parsing)
		parsedUrl, err := url.Parse(movieURL)
		if err != nil {
			fmt.Println("❌ Error parsing URL:", err)
			return
		}
		q := parsedUrl.Query()
		q.Set("apikey", "cad09794")     // Replace with your actual API key
		q.Set("i", idimdb)              // Set the movie ID parameter
		parsedUrl.RawQuery = q.Encode() // saves to parsedUrl.RawQuery
		//send GET request
		response, err := http.Get(parsedUrl.String())
		if err != nil {
			fmt.Println("❌ Error making GET request:", err)
			return
		}
		defer response.Body.Close()

		//read response body
		body, _ := io.ReadAll(response.Body)

		//unmarshal the response body into MovieDetail struct
		var movieDetail MovieDetail
		err = json.Unmarshal(body, &movieDetail)
		if err != nil {
			fmt.Println("❌ Error unmarshalling response:", err)
			return
		}
		//printing the output
		if movieDetail.Response == "True" {
			fmt.Printf("🎬Title: %s\n", movieDetail.Title)
			fmt.Printf("📅Year: %s\n", movieDetail.Year)
			fmt.Printf("⭐Rated: %s\n", movieDetail.Rated)
			fmt.Printf("📅Released: %s\n", movieDetail.Released)
			fmt.Printf("⏳Runtime: %s\n", movieDetail.Runtime)
			fmt.Printf("🎭Genre: %s\n", movieDetail.Genre)
			fmt.Printf("🎬Director: %s\n", movieDetail.Director)
			fmt.Printf("✍️Writer: %s\n", movieDetail.Writer)
			fmt.Printf("👥Actors: %s\n", movieDetail.Actors)
			fmt.Printf("📖Plot: %s\n", movieDetail.Plot)
			fmt.Printf("🌐Language: %s\n", movieDetail.Language)
			fmt.Printf("🌍Country: %s\n", movieDetail.Country)
			fmt.Printf("🏆Awards: %s\n", movieDetail.Awards)
			fmt.Printf("📸Poster: %s\n", movieDetail.Poster)
			fmt.Printf("⭐IMDB Rating: %s\n", movieDetail.IMDBRating)
			fmt.Printf("🔗IMDB ID: %s\n", movieDetail.IMDBID)

		} else {
			fmt.Println("❌ Error:", movieDetail.Error)
		}

	},
}

func init() {
	rootCmd.AddCommand(detailCmd)
	detailCmd.Flags().StringP("id", "i", "", "IMDB ID of the movie to get details for")

}
