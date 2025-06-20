package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestDetailValidIDFields(t *testing.T) {
	baseURL := "https://www.omdbapi.com/"
	parsedUrl, _ := url.Parse(baseURL)
	q := parsedUrl.Query()
	q.Set("apikey", "cad09794")
	q.Set("i", "tt0111161")
	parsedUrl.RawQuery = q.Encode()
	resp, err := http.Get(parsedUrl.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Title      string `json:"Title"`
		Year       string `json:"Year"`
		Plot       string `json:"Plot"`
		Type       string `json:"Type"`
		Poster     string `json:"Poster"`
		IMDBRating string `json:"imdbRating"`
		IMDBID     string `json:"imdbID"`
		Response   string `json:"Response"`
		Error      string `json:"Error"`
	}
	_ = json.Unmarshal(body, &result)
	if result.Response != "True" {
		t.Error("Expected valid movie detail for tt0111161")
	}
	if result.Title == "" || result.Year == "" || result.IMDBID == "" {
		t.Error("Expected non-empty fields for valid movie")
	}
}

func TestDetailInvalidIDErrorMessage(t *testing.T) {
	baseURL := "https://www.omdbapi.com/"
	parsedUrl, _ := url.Parse(baseURL)
	q := parsedUrl.Query()
	q.Set("apikey", "cad09794")
	q.Set("i", "tt0000000")
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
		t.Error("Expected error for invalid ID")
	}
}

func TestDetailMalformedIDErrorMessage(t *testing.T) {
	baseURL := "https://www.omdbapi.com/"
	parsedUrl, _ := url.Parse(baseURL)
	q := parsedUrl.Query()
	q.Set("apikey", "cad09794")
	q.Set("i", "malformed_id")
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
		t.Error("Expected error for malformed ID")
	}
}
