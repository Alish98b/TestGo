package models

import "time"

//new model film
type Film struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Details     string    `json:"details"`
	Genre       string    `json:"genre"`
	Year        time.Time `json:"year"`
}
