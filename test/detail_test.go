package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestOMDbDetailAPI(t *testing.T) {
	baseURL := "https://www.omdbapi.com/"
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		t.Fatal(err)
	}
	q := parsedUrl.Query()
	q.Set("apikey", "cad09794")
	q.Set("i", "tt0111161") // The Shawshank Redemption
	parsedUrl.RawQuery = q.Encode()
	resp, err := http.Get(parsedUrl.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var result struct {
		Title    string `json:"Title"`
		Response string `json:"Response"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}
	if result.Response != "True" {
		t.Errorf("Expected Response True, got %s", result.Response)
	}
	if result.Title != "The Shawshank Redemption" {
		t.Errorf("Expected title The Shawshank Redemption, got %s", result.Title)
	}
}
