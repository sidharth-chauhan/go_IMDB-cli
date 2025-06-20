package models

import "gorm.io/gorm"

// Favorite represents a movie saved as favorite in the local database
type Favorite struct {
	gorm.Model
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDBID string `json:"imdbID" gorm:"unique"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
