package models

type MovieDetail struct {
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
