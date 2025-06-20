package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

type MovieResponse struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
type SearchResult struct {
	Search       []MovieResponse `json:"Search"`
	TotalResults string          `json:"totalResults"`
	Response     string          `json:"Response"`
	Error        string          `json:"Error"` // only present if there's an error
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a movie by title",
	Run: func(cmd *cobra.Command, args []string) {

		// 1️⃣ Read flag from command line
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("❌ Please provide a movie title using --title or -t")
			return
		}
		movieURL := "https://www.omdbapi.com/?apikey=YOUR_KEY&s=THE_TITLE."
		parsedUrl, err := url.Parse(movieURL)
		if err != nil {
			fmt.Println("❌ Error parsing URL:", err)
			return
		}
		updatedTitle := strings.ReplaceAll(title, " ", "+") // Replace spaces with '+' for URL encoding

		q := parsedUrl.Query()
		q.Set("apikey", "cad09794")
		q.Set("s", updatedTitle)        // Set the search query parameter
		parsedUrl.RawQuery = q.Encode() // saves to parsedUrl.RawQuery

		// 5️⃣ Send GET request
		response, err := http.Get(parsedUrl.String())
		if err != nil {
			fmt.Println("❌ Error making GET request:", err)
			return
		}
		defer response.Body.Close()

		//reading the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("❌ Error reading response body:", err)
			return
		}
		//unmarshalling the JSON response
		var result SearchResult
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("❌ Error unmarshalling JSON:", err)
			return
		}
		//printing the search results
		for _, movie := range result.Search {
			fmt.Printf("\n🎬 Title: %s\n📅 Year: %s\n🆔 IMDb ID: %s\n🖼️ Poster: %s\n",
				movie.Title, movie.Year, movie.IMDBID, movie.Poster)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("title", "t", "", "Movie title to search")
}
