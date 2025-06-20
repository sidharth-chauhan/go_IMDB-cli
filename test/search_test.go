package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestOMDbSearchAPI(t *testing.T) {
	baseURL := "https://www.omdbapi.com/"
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		t.Fatal(err)
	}
	q := parsedUrl.Query()
	q.Set("apikey", "cad09794")
	q.Set("s", "Inception")
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
		Response string `json:"Response"`
		Search   []struct {
			Title string `json:"Title"`
		} `json:"Search"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}
	if result.Response != "True" {
		t.Errorf("Expected Response True, got %s", result.Response)
	}
	if len(result.Search) == 0 {
		t.Error("Expected at least one search result")
	}
}
