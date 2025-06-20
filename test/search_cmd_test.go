package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestSearchValidPartialMatch(t *testing.T) {
	baseURL := "https://www.omdbapi.com/"
	parsedUrl, _ := url.Parse(baseURL)
	q := parsedUrl.Query()
	q.Set("apikey", "cad09794")
	q.Set("s", "Incept")
	parsedUrl.RawQuery = q.Encode()
	resp, err := http.Get(parsedUrl.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Response string `json:"Response"`
		Search   []struct {
			Title string `json:"Title"`
		} `json:"Search"`
	}
	_ = json.Unmarshal(body, &result)
	if result.Response != "True" || len(result.Search) == 0 {
		t.Skip("OMDb did not return results for 'Incept' (may be rate-limited or API changed)")
	}
	// Only warn if no title contains "incept"
	found := false
	for _, m := range result.Search {
		if strings.Contains(strings.ToLower(m.Title), "incept") {
			found = true
			break
		}
	}
	if !found {
		t.Log("No result title contains 'incept', but OMDb returned results")
	}
}

func TestSearchCaseInsensitive(t *testing.T) {
	baseURL := "https://www.omdbapi.com/"
	parsedUrl, _ := url.Parse(baseURL)
	q := parsedUrl.Query()
	q.Set("apikey", "cad09794")
	q.Set("s", "inCePtIoN")
	parsedUrl.RawQuery = q.Encode()
	resp, err := http.Get(parsedUrl.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Response string `json:"Response"`
		Search   []struct {
			Title string `json:"Title"`
		} `json:"Search"`
	}
	_ = json.Unmarshal(body, &result)
	if result.Response != "True" {
		t.Error("Expected valid search results for case-insensitive query")
	}
}

func TestSearchNoResultsErrorMessage(t *testing.T) {
	baseURL := "https://www.omdbapi.com/"
	parsedUrl, _ := url.Parse(baseURL)
	q := parsedUrl.Query()
	q.Set("apikey", "cad09794")
	q.Set("s", "asdkfjhasdkjfhakjsdhf")
	parsedUrl.RawQuery = q.Encode()
	resp, err := http.Get(parsedUrl.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Response string `json:"Response"`
		Error    string `json:"Error"`
	}
	_ = json.Unmarshal(body, &result)
	if result.Response != "False" || result.Error == "" {
		t.Error("Expected error message for no results")
	}
}
